package application

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupGin(controller *PostController) *gin.Engine {
	g := gin.New()
	g.GET("/:basePath", controller.Get)
	fs := http.FileServer(http.Dir("assets"))
	g.GET("/:basePath/:filename", func(ctx *gin.Context) {
		http.StripPrefix("/assets", fs).ServeHTTP(ctx.Writer, ctx.Request)
	})
	return g
}
