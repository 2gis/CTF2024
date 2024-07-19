package routes

import (
	"context"
	"fmt"
	auth2 "gopher/auth"
	"gopher/db"
	"gopher/render"
	"gopher/utils"
	"net/http"
	"net/url"
	"strconv"
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
		if len(parts) == 3 {
			id := parts[0]
			title_encoded := parts[1]
			sign := parts[2]
			if utils.Sha512([]byte(fmt.Sprintf("%s_%s_%s", id, title_encoded, auth2.SECRET_KEY))) == sign {
				title_decoded, err := url.QueryUnescape(title_encoded)
				if err != nil {
					return
				}
				key, err := auth2.KeyStr(id, title_decoded)
				if err != nil {
					return
				}
				atoi, _ := strconv.Atoi(key)
				var title, content string
				err = db.Instance.QueryRow(context.Background(), "SELECT title, content FROM mails WHERE hkey = $1;", atoi).Scan(&title, &content)
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