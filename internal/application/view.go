package application

import (
	"html/template"
	"net/http"
)

type View struct {
	TemplatePath string
	Data         interface{}
}

func (v *View) Render(w http.ResponseWriter) error {
	t, err := template.ParseFiles("templates/post.tmpl")
	if err != nil {
		return err
	}
	if err := t.Execute(w, v.Data); err != nil {
		return err
	}
	return nil
}

func (v *View) WriteContentType(w http.ResponseWriter) {
	w.Header()["Content-Type"] = []string{"text/html"}
}
