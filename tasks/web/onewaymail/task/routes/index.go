package routes

import (
	"context"
	"fmt"
	auth2 "gopher/auth"
	"gopher/db"
	"gopher/render"
	"gopher/utils"
	"net/http"
	"strings"
)

func Index(w http.ResponseWriter, r *http.Request) {
	letter := ""
	auth := ""
	for _, cookie := range r.Cookies() {
		if cookie.Name == "auth" {
			auth = cookie.Value
		}
	}
	if auth != "" {
		parts := strings.Split(auth, "_")
		if len(parts) == 2 {
			id := parts[0]
			sign := parts[1]
			if utils.Sha512([]byte(fmt.Sprintf("%s_%s", id, auth2.SECRET_KEY))) == sign {
				var title, content string
				err := db.Instance.QueryRow(context.Background(), "SELECT title, content FROM mails WHERE id = $1;", id).Scan(&title, &content)
				if err != nil {
					if err.Error() != "no rows in result set" {
						fmt.Println(err.Error())
						return
					}
				}
				letter = "Your last mail: " + title + "<br>" + content
			}
		}
	}
	w.Write([]byte(render.GetView("index").Render(map[string]string{
		"letter": letter,
	})))
}