package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/keitam913/textvid/application/blog"
	"github.com/keitam913/textvid/domain"
)

type PostController struct {
	Service        blog.Service
	ViewRepository *ViewRepository
}

func (pc *PostController) Get(ctx *gin.Context) {
	fn := ctx.Param("basePath")
	post, err := pc.Service.GetByHTMLFilename(fn)
	if _, ok := err.(domain.ErrNotFound); ok {
		ctx.AbortWithError(http.StatusNotFound, err)
		return
	}
	if err != nil {
		panic(err)
	}
	ctx.Render(http.StatusOK, pc.ViewRepository.Post(post))
}

func (pc *PostController) List(ctx *gin.Context) {
	posts, err := pc.Service.Recent()
	if err != nil {
		panic(err)
	}
	ctx.Render(http.StatusOK, pc.ViewRepository.List(posts))
}

type WebhookController struct {
	Service blog.Service
}

func (whc *WebhookController) Post(ctx *gin.Context) {
	var ev domain.PushEvent
	if err := ctx.Bind(&ev); err != nil {
		panic(err)
	}
	if err := whc.Service.PushPosts(&ev); err != nil {
		panic(err)
	}
}
