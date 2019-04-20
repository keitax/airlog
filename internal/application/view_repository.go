package application

import "github.com/keitax/airlog/internal/domain"

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
