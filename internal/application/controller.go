package application

import (
	"github.com/gin-gonic/gin"
	"github.com/keitax/airlog/internal/domain"
	"log"
	"net/http"
)

type PostController struct {
	Service        domain.PostService
	ViewRepository *ViewRepository
}

func (pc *PostController) Get(ctx *gin.Context) {
	fn := ctx.Param("basePath")
	post, err := pc.Service.GetByHTMLFilename(fn)
	if _, ok := err.(domain.ErrNotFound); ok {
		ctx.AbortWithError(http.StatusNotFound, err)
		return
	} else if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	log.Printf("title %s", post.Title)

	ctx.Render(http.StatusOK, pc.ViewRepository.Post(post))
}
