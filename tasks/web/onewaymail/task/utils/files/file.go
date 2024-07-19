package files

import (
	"errors"
	"io"
	"os"
)

type File struct {
	Path string
	file *os.File
}

func (file *File) Open(flag int) error {
	if file.Path == "" {
		err := errors.New("path not specified.")
		return err
	}
	var err error
	file.file, err = os.OpenFile(file.Path, flag, 0)
	if err != nil {
		return err
	}
	return nil
}

func (file *File) Create() error {
	if file.Path == "" {
		err := errors.New("path not specified.")
		return err
	}
	var err error
	file.file, err = os.Create(file.Path)
	if err != nil {
		return err
	}
	return nil
}

func (file *File) Close() error {
	err := file.file.Close()
	if err != nil {
		return err
	}
	return nil
}

func (file *File) Read() []byte {
	content := []byte{}
	stat, _ := file.file.Stat()
	data := make([]byte, stat.Size())
	for {
		n, err := file.file.Read(data)
		if err == io.EOF {
			break
		}
		content = append(content, data[:n]...)
	}
	return content
}

func (file *File) ReadString() string {
	return string(file.Read())
}

func (file *File) Exists() bool {
	if file.Path == "" {
		return false
	}
	_, err := os.Stat(file.Path)
	return !os.IsNotExist(err)
}

func (file *File) Write(content []byte) error {
	_, err := file.file.Write(content)
	return err
}

func (file *File) WriteString(content string) error {
	_, err := file.file.WriteString(content)
	return err
}

func (file *File) Remove() error {
	return os.RemoveAll(file.Path)
}

func (file *File) IsFile() bool {
	obj, err := os.Stat(file.Path)
	if err != nil {
		return false
	}
	return !obj.IsDir()
}

func (file *File) IsDir() bool {
	obj, err := os.Stat(file.Path)
	if err != nil {
		return false
	}
	return obj.IsDir()
}
