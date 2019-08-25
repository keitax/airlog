package apigatewayproxy

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

type Handler struct {
}

func (h *Handler) Handle(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       "test response",
	}, nil
}
