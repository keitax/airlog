package web

import (
	"fmt"
	"github.com/keitam913/airlog/domain"
	"github.com/russross/blackfriday"
	"html/template"
	"net/http"
	"path/filepath"
	"strings"
	"time"
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
		"GetPostURL":    GetPostURL,
		"ParseMarkdown": ParseMarkdown,
		"ShowDate":      ShowDate,
		"ShowLabels": func(p *domain.Post) string {
			return strings.Join(p.Labels, ", ")
		},
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

func ParseMarkdown(text string) template.HTML {
	bs := blackfriday.Run([]byte(text))
	return template.HTML(string(bs))
}

func ShowDate(t time.Time) string {
	return t.Format("2006-01-02")
}
