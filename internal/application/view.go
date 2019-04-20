package application

import (
	"github.com/keitax/airlog/internal/domain"
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

type ViewRepository struct {
	SiteTitle string
}

func (v *ViewRepository) Post(post *domain.Post) *View {
	return &View{
		TemplatePath: "templates/post.tmpl",
		Data: map[string]interface{}{
			"siteTitle": v.SiteTitle,
			"post":      post,
		},
	}
}
