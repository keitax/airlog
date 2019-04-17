package di

import (
	"github.com/gin-gonic/gin"
	"github.com/keitax/airlog/internal/application"
	"github.com/keitax/airlog/internal/domain"
	"github.com/keitax/airlog/internal/infrastructure/osenv"
)

type Container struct{}

func (c Container) Gin() *gin.Engine {
	g := application.SetupGin(c.PostController())
	g.Use(gin.Recovery(), gin.Logger())
	return g
}

func (c Container) PostController() *application.PostController {
	return &application.PostController{
		Service: c.PostService(),
		View:    c.View(),
	}
}

func (c Container) PostService() domain.PostService {
	return &domain.PostServiceImpl{}
}

func (c Container) View() *application.View {
	return &application.View{}
}

func (c Container) Config() *application.Config {
	conf, err := osenv.LoadConfig()
	if err != nil {
		panic(err)
	}
	return conf
}
