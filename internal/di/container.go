package di

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
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
		ViewRepository: c.ViewRepository(),
	}
}

func (c Container) PostService() domain.PostService {
	return &domain.PostServiceImpl{
		Repository: c.PostRepository(),
	}
}

func (c Container) PostRepository() domain.PostRepository {
	return &rds.PostRepository{
		DB: c.DB(),
	}
}

func (c Container) ViewRepository() *application.ViewRepository {
	return &application.ViewRepository{
		SiteTitle: c.Config().SiteTitle,
		Footnote:  c.Config().Footnote,
	}
}

func (c Container) DB() *sql.DB {
	db, err := sql.Open("mysql", c.Config().BlogDSN)
	if err != nil {
		panic(err)
	}
	return db
}

func (c Container) Config() *application.Config {
	conf, err := osenv.LoadConfig()
	if err != nil {
		panic(err)
	}
	return conf
}
