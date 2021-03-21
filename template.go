package main

import (
	"html/template"
	"log"
	"net/url"

	"github.com/forewing/webrcon-server/version"
)

const (
	queryHashLength = 7
)

func mustLoadTemplate() *template.Template {
	t, err := template.New("").Delims("[[", "]]").Funcs(template.FuncMap{
		"generateStaticURL": func(origin string) string {
			u, err := url.Parse(origin)
			if err != nil {
				panic(err)
			}
			if version.Hash != version.HashDefault && len(version.Hash) >= queryHashLength {
				q := u.Query()
				q.Set("v", version.Hash[0:queryHashLength])
				u.RawQuery = q.Encode()
			}
			return u.String()
		},
	}).ParseFS(templates, "*.html")

	if err != nil {
		log.Panicln(err)
	}
	return t
}
