package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupGin(controller *PostController, webhookController *WebhookController) *gin.Engine {
	g := gin.New()

	g.GET("/", controller.List)
	g.GET("/:basePath", func(ctx *gin.Context) {
		if ctx.Param("basePath") == "health" {
			ctx.String(http.StatusOK, "OK")
			return
		}
		controller.Get(ctx)
	})
	fs := http.FileServer(http.Dir("assets"))
	g.GET("/:basePath/:filename", func(ctx *gin.Context) {
		http.StripPrefix("/assets", fs).ServeHTTP(ctx.Writer, ctx.Request)
	})

	g.POST("/webhook", webhookController.Post)

	return g
}
