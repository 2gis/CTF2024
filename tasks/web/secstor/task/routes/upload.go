package routes

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

func Upload(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)
	u, err := url.Parse(r.RequestURI)
	if err != nil {
		return
	}
	file, handler, err := r.FormFile("file")
	if err != nil {
		return
	}
	var filename string
	if u.Query().Has("filename") {
		filename = u.Query().Get("filename")
	} else {
		filename = handler.Filename
	}
	defer file.Close()
	dst, err := os.Create("./static/" + filename)
	defer dst.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte(fmt.Sprintf("Successfully Uploaded File. Link: /static/%s", filename)))
}