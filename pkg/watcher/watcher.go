package watcher // import "github.com/muhhae/go-templ-htmx/pkg/watcher"

import (
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/fsnotify/fsnotify"
)

type WatchConfig struct {
	Command     []string `json:"command"`
	IncludeDirs []string `json:"includeDirs"`
	ExcludeDirs []string `json:"excludeDirs"`
	Exclude     []string `json:"exclude"`
	Include     []string `json:"include"`
}

func (c WatchConfig) Run() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	eventTime := make(map[string]time.Time)

	for _, dir := range c.IncludeDirs {
		err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				color.Red(err.Error())
				return nil
			}
			if info.IsDir() {
				for _, s := range c.ExcludeDirs {
					if info.Name() == s {
						color.Red("excluding " + path)
						return filepath.SkipDir
					}
				}
				err = watcher.Add(path)
				if err != nil {
					color.Red(err.Error())
				} else {
					color.Yellow("watching " + path)
				}
			}
			return nil
		})
		if err != nil {
			color.Red(err.Error())
		}
	}
	color.Green("\nwatching...\n")

	go func() {
		command(c.Command)
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if t, ok := eventTime[event.Name]; ok && time.Since(t) < 500*time.Millisecond {
					continue
				}
				eventTime[event.Name] = time.Now()
				if event.Op&fsnotify.Write == fsnotify.Write {
					if !c.includeCheck(event.Name) {
						continue
					}
					fileName := path.Base(event.Name)
					color.Yellow("modified " + fileName)
					color.Green("rebuilding...")
					go command(c.Command)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				color.Red(err.Error())
			}
		}
	}()

	<-done
}

var cmd *exec.Cmd

func command(l []string) {
	if cmd != nil {
		cmd.Process.Kill()
		cmd.Wait()
	}

	for i, s := range l {
		color.Yellow("running " + s)
		parts := strings.Fields(s)
		cmd = exec.Command(parts[0], parts[1:]...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Start()
		if err != nil {
			color.Red(err.Error())
			color.Red("rebuilding halted...")
			return
		}
		if i == len(l)-1 {
			color.Green("rebuilding complete")
			return
		}
		err = cmd.Wait()
		if err != nil {
			color.Red(err.Error())
			color.Red("rebuilding halted...")
			return
		}
	}
}

func (c WatchConfig) includeCheck(f string) bool {
	for _, s := range c.Exclude {
		matched, err := path.Match(s, f)
		if err != nil {
			color.Red(err.Error())
			continue
		}
		if matched {
			return false
		}
	}

	for _, s := range c.Include {
		matched, err := path.Match(s, f)
		if err != nil {
			color.Red(err.Error())
			continue
		}
		if matched {
			return true
		}
	}

	return false
}
