package application

import "github.com/keitax/airlog/internal/domain"

type View interface {
	Render(post *domain.Post) ([]byte, error)
}
