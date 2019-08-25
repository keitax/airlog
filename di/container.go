package di

import (
	"database/sql"

	"github.com/keitam913/airlog/application/blog"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/keitam913/airlog/domain"
	"github.com/keitam913/airlog/infrastructure/ghapi"
	"github.com/keitam913/airlog/infrastructure/osenv"
	"github.com/keitam913/airlog/infrastructure/rds"
	"github.com/keitam913/airlog/infrastructure/web"
)

type Container struct{}

func (c Container) Gin() *gin.Engine {
	g := web.SetupGin(c.PostController(), c.WebhookController())
	g.Use(gin.Recovery(), gin.Logger())
	return g
}

func (c Container) PostController() *web.PostController {
	return &web.PostController{
		Service:        c.BlogService(),
		ViewRepository: c.ViewRepository(),
	}
}

func (c Container) WebhookController() *web.WebhookController {
	return &web.WebhookController{
		Service:          c.BlogService(),
		PostFileRepository: c.PostFileRepository(),
	}
}

func (c Container) BlogService() blog.Service {
	return &blog.ServiceImpl{
		Service:    c.PostService(),
		Repository: c.PostRepository(),
	}
}

func (c Container) PostService() domain.PostService {
	return &domain.PostServiceImpl{}
}

func (c Container) PostRepository() domain.PostRepository {
	return &rds.PostRepository{
		DB: c.DB(),
	}
}

func (c Container) ViewRepository() *web.ViewRepository {
	return &web.ViewRepository{
		SiteTitle: c.Config().SiteTitle,
		Footnote:  c.Config().Footnote,
	}
}

func (c Container) PostFileRepository() domain.PostFileRepository {
	return &ghapi.PostFileRepository{
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

func (c Container) Config() *osenv.Config {
	conf, err := osenv.LoadConfig()
	if err != nil {
		panic(err)
	}
	return conf
}
