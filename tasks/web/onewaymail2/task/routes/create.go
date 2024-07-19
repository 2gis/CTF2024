package routes

import (
	"context"
	"fmt"
	auth2 "gopher/auth"
	"gopher/db"
	"gopher/utils"
	"net/http"
	"net/url"
	"strconv"
)

func Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		return
	}
	if err := r.ParseForm(); err != nil {
		fmt.Println(err.Error())
		return
	}
	title_decoded := r.Form.Get("title")
	title := url.QueryEscape(title_decoded)
	content := r.Form.Get("content")
	var id int
	err := db.Instance.QueryRow(context.Background(), "INSERT INTO mails (title, content) VALUES ($1, $2) RETURNING id;", title, content).Scan(&id)
	if err != nil {
		if err.Error() != "no rows in result set" {
			fmt.Println(err.Error())
			return
		}
	}
	key, err := auth2.KeyStr(strconv.Itoa(id), title_decoded)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	atoi, _ := strconv.Atoi(key)
	err = db.Instance.QueryRow(context.Background(), "UPDATE mails SET hkey = $1 WHERE id = $2;", atoi, id).Scan()
	if err != nil {
		if err.Error() != "no rows in result set" {
			fmt.Println(err.Error())
			return
		}
	}
	cookie := &http.Cookie{Name: "auth", Value: fmt.Sprintf("%d_%s_%s", id, title, utils.Sha512([]byte(fmt.Sprintf("%d_%s_%s", id, title, auth2.SECRET_KEY))))}
	http.SetCookie(w, cookie)
	w.Write([]byte("mail sent"))
}