package application

import "github.com/keitax/airlog/internal/domain"

type PostController struct {
	Service domain.PostService
	View    View
}
