package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/muhhae/go-templ-htmx/pkg/watcher"
)

func main() {
	args := os.Args[1:]
	c := watcher.WatchConfig{}
	if len(args) == 0 {
		c_json, err := os.ReadFile("watcher.json")
		if err != nil {
			fmt.Println("no watcher.json file found")
			return
		}
		err = json.Unmarshal(c_json, &c)
		if err != nil {
			fmt.Println(err)
			return
		}
	} else {
		if args[0] == "init" {
			c = watcher.WatchConfig{
				Command:     []string{},
				IncludeDirs: []string{},
				ExcludeDirs: []string{},
				Exclude:     []string{},
				Include:     []string{},
			}
			c_json, err := json.MarshalIndent(c, "", "  ")
			if err != nil {
				fmt.Println(err)
				return
			}
			if _, err := os.Stat("watcher.json"); err == nil {
				fmt.Println("watcher.json already exists")
				return
			}
			err = os.WriteFile("watcher.json", c_json, 0644)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("watcher.json created")
		} else {
			for i := 0; i < len(args); i++ {
				if args[i] == "-c" {
					i++
					c_json, err := os.ReadFile(args[i])
					if err != nil {
						fmt.Println(err)
						return
					}
					err = json.Unmarshal(c_json, &c)
					if err != nil {
						fmt.Println(err)
						return
					}
					continue
				}
				fmt.Println("unknown argument: " + args[i])
				return
			}
		}
	}
	c.Run()
}
