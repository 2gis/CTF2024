package routes

import (
	"gopher/plugin"
	"net/http"
)

func Reload(w http.ResponseWriter, r *http.Request) {
	plugin.DisableAllPlugins()
	plugin.EnableAllPlugins()
	w.Write([]byte("Reload done!"))
}