package render

import (
	html2 "html"
	"regexp"
	"strings"
)

type GHtml string

func (html GHtml) Render(data map[string]string) string {
	render := string(html)
	allInserts := regexp.MustCompile("<%[&=?!-][ a-zA-Z0-9\\/\\-_]+%>").FindAllString(string(html), -1)
	for _, insert := range allInserts {
		insert_tmp := insert
		key, val := "", ""
		insert = strings.ReplaceAll(insert, " ", "")
		for _, char := range []rune(insert) {
			if string(char) == "=" {
				key = insert[3 : len(insert)-2]
				val = data[key]
			}else if string(char) == "!" {
				key = insert[3 : len(insert)-2]
				val = html2.EscapeString(data[key])
			} else if string(char) == "&" {
				key = insert[3 : len(insert)-2]
				val = GetView(key).Render(data)
			}
		}
		render = strings.ReplaceAll(render, insert_tmp, val)
	}
	return render
}
