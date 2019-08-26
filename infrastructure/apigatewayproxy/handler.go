package apigatewayproxy

import (
	"context"

	"github.com/gin-gonic/gin"

	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"

	"github.com/aws/aws-lambda-go/events"
)

type Handler struct {
	ginLambda *ginadapter.GinLambda
}

func NewHandler(eng *gin.Engine) *Handler {
	return &Handler{
		ginLambda: ginadapter.New(eng),
	}
}

func (h *Handler) Handle(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return h.ginLambda.ProxyWithContext(ctx, req)
}
