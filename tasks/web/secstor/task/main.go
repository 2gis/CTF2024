package main

import (
	"gopher/plugin"
	"gopher/render"
	"gopher/routes"
	"net/http"
)

func main() {
	http.HandleFunc("/", routes.Index)
	http.HandleFunc("/reload", routes.Reload)
	http.HandleFunc("/upload", routes.Upload)

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	render.LoadTemplates()
	plugin.EnableAllPlugins()
	http.ListenAndServe(":5005", nil)
}