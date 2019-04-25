package application

import (
	"fmt"
	"github.com/keitax/airlog/internal/domain"
	"html/template"
	"net/http"
	"path/filepath"
	"strings"
)

type View struct {
	TemplatePath string
	Data         interface{}
}

func (v *View) Render(w http.ResponseWriter) error {
	fs, err := filepath.Glob("templates/*.tmpl")
	if err != nil {
		return err
	}
	t, err := template.New("root").Funcs(template.FuncMap{
		"GetPostURL": GetPostURL,
	}).ParseFiles(fs...)
	if err != nil {
		return err
	}
	if err := t.ExecuteTemplate(w, strings.Replace(v.TemplatePath, "templates/", "", -1), v.Data); err != nil {
		return err
	}
	return nil
}

func (v *View) WriteContentType(w http.ResponseWriter) {
	w.Header()["Content-Type"] = []string{"text/html"}
}

func GetPostURL(post *domain.Post) string {
	return fmt.Sprintf("/%s", strings.Replace(post.Filename, ".md", ".html", -1))
}
