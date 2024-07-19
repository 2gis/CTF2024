package routes

import (
	"gopher/render"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(render.GetView("index").Render(map[string]string{})))
}