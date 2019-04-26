package application

import (
	"github.com/keitax/airlog/internal/domain"
)

type ViewRepository struct {
	SiteTitle string
	Footnote  string
}

func (v *ViewRepository) Post(post *domain.Post) *View {
	return &View{
		TemplatePath: "templates/post.tmpl",
		Data: map[string]interface{}{
			"siteTitle": v.SiteTitle,
			"post":      post,
			"footnote":  v.Footnote,
		},
	}
}

func (v *ViewRepository) List(posts []*domain.Post) *View {
	return &View{
		TemplatePath: "templates/list.tmpl",
		Data: map[string]interface{}{
			"siteTitle": v.SiteTitle,
			"posts":     posts,
			"footnote":  v.Footnote,
		},
	}
}
