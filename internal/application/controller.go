package application

import (
	"github.com/gin-gonic/gin"
	"github.com/keitax/airlog/internal/domain"
	"net/http"
)

type PostController struct {
	Service domain.PostService
	View    *View
}

func (pc *PostController) Get(ctx *gin.Context) {
	fn := ctx.Param("filename")
	post, err := pc.Service.GetByHTMLFilename(fn)
	if _, ok := err.(domain.ErrNotFound); ok {
		ctx.AbortWithError(http.StatusNotFound, err)
		return
	} else if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	html, err := pc.View.RenderPost(post)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.String(http.StatusOK, html)
}
