package routes

import (
	"context"
	"fmt"
	auth2 "gopher/auth"
	"gopher/db"
	"gopher/utils"
	"net/http"
	"strings"
)

func Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		return
	}
	if err := r.ParseForm(); err != nil {
		fmt.Println(err)
		return
	}
	title := r.Form.Get("title")
	content := r.Form.Get("content")
	title = strings.ReplaceAll(title, "|", "")
	title = strings.ReplaceAll(title, "*", "")
	title = strings.ReplaceAll(title, "/", "")
	title = strings.ReplaceAll(title, "&", "")
	content = strings.ReplaceAll(content, "|", "")
	content = strings.ReplaceAll(content, "*", "")
	content = strings.ReplaceAll(content, "/", "")
	content = strings.ReplaceAll(content, "&", "")
	var id int
	err := db.Instance.QueryRow(context.Background(), fmt.Sprintf("INSERT INTO mails (title, content) VALUES ('%s', '%s') RETURNING id;", title, content)).Scan(&id)
	if err != nil {
		if err.Error() != "no rows in result set" {
			return
		}
	}
	cookie := &http.Cookie{Name: "auth", Value: fmt.Sprintf("%d_%s", id, utils.Sha512([]byte(fmt.Sprintf("%d_%s", id, auth2.SECRET_KEY))))}
	http.SetCookie(w, cookie)
	w.Write([]byte("mail sent"))
}