package db

import (
	"context"
	"gopher/utils/files"
	"os"
	"strings"
)

func ExecuteSQLFile(file files.File) error {
	err := file.Open(os.O_RDWR)
	if err != nil {
		return err
	}
	content := file.ReadString()
	for _, cmd := range strings.Split(content, ";") {
		_, err = Instance.Exec(context.Background(), cmd)
		if err != nil {
			return err
		}
	}
	return nil
}
