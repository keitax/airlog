package di

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	awsdynamodb "github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/keitam913/textvid/infrastructure/apigatewayproxy"
	"github.com/keitam913/textvid/infrastructure/dynamodb"

	"github.com/keitam913/textvid/application/blog"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/keitam913/textvid/domain"
	"github.com/keitam913/textvid/infrastructure/ghapi"
	"github.com/keitam913/textvid/infrastructure/osenv"
	"github.com/keitam913/textvid/infrastructure/web"
)

type Container struct{}

func (c Container) APIGatewayProxyHandler() *apigatewayproxy.Handler {
	return apigatewayproxy.NewHandler(c.Gin())
}

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
		Service:            c.BlogService(),
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
	conf := c.Config()
	if conf.Mode == "local" {
		return &dynamodb.PostRepository{
			DB: c.DynamoDBLocal(),
		}
	}
	return &dynamodb.PostRepository{
		DB: c.DynamoDB(),
	}
}

func (c Container) DynamoDB() *awsdynamodb.DynamoDB {
	return awsdynamodb.New(session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	})))
}

func (c Container) DynamoDBLocal() *awsdynamodb.DynamoDB {
	return awsdynamodb.New(session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region:   aws.String("local"),
			Endpoint: aws.String("http://dynamodb:8000"),
		},
	})))
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

func (c Container) Config() *osenv.Config {
	conf, err := osenv.LoadConfig()
	if err != nil {
		panic(err)
	}
	return conf
}
