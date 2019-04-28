package application

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/keitax/airlog/internal/domain"
	"io/ioutil"
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
	ctx.Render(http.StatusOK, pc.ViewRepository.Post(post))
}

func (pc *PostController) List(ctx *gin.Context) {
	posts, err := pc.Service.Recent()
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.Render(http.StatusOK, pc.ViewRepository.List(posts))
}

type WebhookController struct {
	PostService      domain.PostService
	GitHubRepository domain.GitHubRepository
}

func (whc *WebhookController) Post(ctx *gin.Context) {
	bs, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	var ev domain.PushEvent
	if err := json.Unmarshal(bs, &ev); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	fs, err := whc.GitHubRepository.ChangedFiles(&ev)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	for _, f := range fs {
		if domain.IsPostFileName(f.Path) {
			if err := whc.PostService.RegisterPost(f.Path, f.Content); err != nil {
				ctx.AbortWithError(http.StatusInternalServerError, err)
				return
			}
		}
	}
}
