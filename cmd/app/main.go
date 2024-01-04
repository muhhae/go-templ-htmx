package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/muhhae/go-templ-htmx/internal/view"
)

func main() {
	component := view.Hello("World")
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("internal/static"))))

	http.Handle("/", templ.Handler(component))

	fmt.Println("Listening on http://localhost:3000")
	fmt.Println("Press Ctrl+C to stop the web server...")
	http.ListenAndServe("localhost:3000", nil)
}
