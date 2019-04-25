package application

import (
	"github.com/keitax/airlog/internal/domain"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
	"html/template"
)

type ViewRepository struct {
	SiteTitle string
}

func (v *ViewRepository) Post(post *domain.Post) *View {
	return &View{
		TemplatePath: "templates/post.tmpl",
		Data: map[string]interface{}{
			"siteTitle":    v.SiteTitle,
			"post":         post,
			"renderedBody": ParseMarkdown(post.Body),
		},
	}
}

func (v *ViewRepository) List(posts []*domain.Post) *View {
	return &View{
		TemplatePath: "templates/list.tmpl",
		Data: map[string]interface{}{
			"siteTitle": v.SiteTitle,
			"posts":     posts,
		},
	}
}

func ParseMarkdown(text string) template.HTML {
	bs := blackfriday.Run([]byte(text))
	bs = bluemonday.UGCPolicy().SanitizeBytes(bs)
	return template.HTML(string(bs))
}
