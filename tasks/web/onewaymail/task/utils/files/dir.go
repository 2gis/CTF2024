package files

import (
	"os"
)

type Directory struct {
	Path string
}

func (dir *Directory) Create() error {
	err := os.Mkdir(dir.Path, 0777)
	if err != nil {
		return err
	}
	return nil
}

func (dir *Directory) CreateAll() error {
	err := os.MkdirAll(dir.Path, 0777)
	if err != nil {
		return err
	}
	return nil
}

func (dir *Directory) Remove(path string) error {
	err := os.RemoveAll(path)
	if err != nil {
		return err
	}
	return nil
}
