package di

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/keitam913/airlog/internal/application"
	"github.com/keitam913/airlog/internal/domain"
	"github.com/keitam913/airlog/internal/infrastructure/ghapi"
	"github.com/keitam913/airlog/internal/infrastructure/osenv"
	"github.com/keitam913/airlog/internal/infrastructure/rds"
)

type Container struct{}

func (c Container) Gin() *gin.Engine {
	g := application.SetupGin(c.PostController(), c.WebhookController())
	g.Use(gin.Recovery(), gin.Logger())
	return g
}

func (c Container) PostController() *application.PostController {
	return &application.PostController{
		Service:        c.PostService(),
		ViewRepository: c.ViewRepository(),
	}
}

func (c Container) WebhookController() *application.WebhookController {
	return &application.WebhookController{
		PostService:      c.PostService(),
		GitHubRepository: c.GitHubRepository(),
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

func (c Container) GitHubRepository() domain.GitHubRepository {
	return &ghapi.GitHubRepository{
		GitHubAPIPostRepositoryEndpoint: c.Config().GitHubAPIPostRepositoryEndpoint,
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
