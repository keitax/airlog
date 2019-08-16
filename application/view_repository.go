package application

import (
	"fmt"
	"github.com/keitam913/airlog/domain"
)

type ViewRepository struct {
	SiteTitle string
	Footnote  string
}

func (v *ViewRepository) Post(post *domain.Post) *View {
	return &View{
		TemplatePath: "templates/post.tmpl",
		Data: map[string]interface{}{
			"headerTitle": fmt.Sprintf("%s - %s", post.Title, v.SiteTitle),
			"siteTitle":   v.SiteTitle,
			"post":        post,
			"footnote":    v.Footnote,
		},
	}
}

func (v *ViewRepository) List(posts []*domain.Post) *View {
	return &View{
		TemplatePath: "templates/list.tmpl",
		Data: map[string]interface{}{
			"headerTitle": v.SiteTitle,
			"siteTitle":   v.SiteTitle,
			"posts":       posts,
			"footnote":    v.Footnote,
		},
	}
}
