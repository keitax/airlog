package di

import (
	"github.com/gin-gonic/gin"
	"github.com/keitax/airlog/internal/application"
	"github.com/keitax/airlog/internal/domain"
	"github.com/keitax/airlog/internal/infrastructure/osenv"
	"github.com/keitax/airlog/internal/infrastructure/rds"
)

type Container struct{}

func (c Container) Gin() *gin.Engine {
	g := application.SetupGin(c.PostController())
	g.Use(gin.Recovery(), gin.Logger())
	return g
}

func (c Container) PostController() *application.PostController {
	return &application.PostController{
		Service:        c.PostService(),
		ViewRepository: c.View(),
	}
}

func (c Container) PostService() domain.PostService {
	return &domain.PostServiceImpl{
		Repository: c.PostRepository(),
	}
}

func (c Container) PostRepository() domain.PostRepository {
	return &rds.PostRepository{}
}

func (c Container) View() *application.ViewRepository {
	return &application.ViewRepository{}
}

func (c Container) Config() *application.Config {
	conf, err := osenv.LoadConfig()
	if err != nil {
		panic(err)
	}
	return conf
}
