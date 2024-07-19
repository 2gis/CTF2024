package main

import (
	"gopher/db"
	"gopher/render"
	"gopher/routes"
	"net/http"
)

func main() {
	db.Init()
	http.HandleFunc("/", routes.Index)
	http.HandleFunc("/create", routes.Create)
	render.LoadTemplates()
	http.ListenAndServe(":5007", nil)
}