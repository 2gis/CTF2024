package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"gopher/utils/files"
	"os"
	"strings"
	"time"
)

func Init() {
	conn, err := pgxpool.New(context.Background(), fmt.Sprintf("postgres://%s:%s@%s/%s",
		strings.ReplaceAll(os.Getenv("POSTGRES_USER"), "\r", ""),
		strings.ReplaceAll(os.Getenv("POSTGRES_PASSWORD"), "\r", ""),
		strings.ReplaceAll(os.Getenv("POSTGRES_HOST"), "\r", ""),
		strings.ReplaceAll(os.Getenv("POSTGRES_DB"), "\r", "")))
	if err != nil {
		return
	}
	Instance = conn
	attempts := 0
	for true {
		err = ExecuteSQLFile(files.File{Path: "sql/init_tables.sql"})
		if attempts == 5 && err != nil {
			return
		} else if err != nil {
			attempts++
			time.Sleep(time.Second * 5)
		} else {
			break
		}
	}
}
