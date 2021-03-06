package main

import (
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/keitam913/textvid/di"
)

func main() {
	dc := di.Container{}
	lambda.Start(dc.APIGatewayProxyHandler().Handle)
}
