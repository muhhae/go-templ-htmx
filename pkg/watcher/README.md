# Prerequisities 
1. Go (latest)

# Usage
1. Import the package "github.com/muhhae/go-templ-htmx/pkg/watcher"
2. Create WatchConfig struct with the following fields:
    - Command - Command to run when a change is detected
    - IncludeDirs - List of directory to watch for changes
    - ExcludeDirs - List of directory to ignore
    - Exclude - List of filename or file extension to ignore
    - Include - List of filename or file extension to watch for changes
    <br>Example :
    ```go
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
    ``` 
2. Call c.Run()
3. Execute the program "go run /path/to/main.go"