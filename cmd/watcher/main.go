package main

import "github.com/muhhae/go-templ-htmx/pkg/watcher"

func main() {
	c := watcher.WatchConfig{
		Command: []string{
			"npx tailwindcss -i ./internal/style/style.css -o ./internal/static/style/output.css",
			"templ generate",
			"go run ./cmd/app/main.go",
		},
		IncludeDirs: []string{
			".",
		},
		ExcludeDirs: []string{
			"internal/static",
			"node_modules",
			".git",
			".idea",
			".vscode",
			"vendor",
			"cmd/watcher",
		},
		Exclude: []string{
			"*_templ.go",
			"*_test.go",
			"*.git",
		},
		Include: []string{
			"*.go",
			"*.html",
			"*.css",
			"*.js",
			"*.templ",
		},
	}
	c.Run()
}
