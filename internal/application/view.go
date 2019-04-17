package application

import (
	"bytes"
	"github.com/keitax/airlog/internal/domain"
	"html/template"
)

type View struct{}

func (v *View) RenderPost(post *domain.Post) (string, error) {
	t, err := template.ParseFiles("templates/post.tmpl")
	if err != nil {
		return "", err
	}
	buf := &bytes.Buffer{}
	if err := t.Execute(buf, map[string]string{}); err != nil {
		return "", err
	}
	return buf.String(), nil
}
