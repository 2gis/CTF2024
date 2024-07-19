package render

import (
	"gopher/utils/files"
	"io/ioutil"
	"os"
	"path/filepath"
)

var templates_storage map[string]GHtml

func appendFiles(arr []string, path string) []string {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return []string{}
	}

	for _, dir := range files {
		if dir.IsDir() {
			arr = appendFiles(arr, path+dir.Name()+"/")
		}
	}
	to_append, err := filepath.Glob(path + "*.html")
	if err != nil {
		return []string{}
	}
	arr = append(arr, to_append...)
	return arr
}

func LoadTemplates() {
	templates_storage = make(map[string]GHtml)
	viewsDir := files.Directory{Path: "views"}
	err := viewsDir.CreateAll()
	if err != nil {
		return
	}

	all_files := appendFiles([]string{}, "views/")
	for _, filename := range all_files {
		file := files.File{Path: filename}
		err := file.Open(os.O_RDWR)
		if err != nil {
			return
		}
		content := GHtml(file.ReadString())
		err = file.Close()
		if err != nil {
			return
		}
		templates_storage[filename[6:len(filename)-5]] = content
	}
}

func GetView(template string) GHtml {
	return templates_storage[template]
}
