package application

import "github.com/gin-gonic/gin"

func SetupGin(controller *PostController) *gin.Engine {
	g := gin.New()
	g.GET("/:filename", controller.Get)
	return g
}
