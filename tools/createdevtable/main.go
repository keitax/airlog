package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {
	db := dynamodb.New(session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region:   aws.String("local"),
			Endpoint: aws.String("http://localhost:8000"),
		},
	})))

	if _, err := db.CreateTable(&dynamodb.CreateTableInput{
		TableName: aws.String("Post"),
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("Filename"),
				KeyType:       aws.String(dynamodb.KeyTypeHash),
			},
		},
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("Filename"),
				AttributeType: aws.String(dynamodb.ScalarAttributeTypeS),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(5),
			WriteCapacityUnits: aws.Int64(5),
		},
	}); err != nil {
		panic(err)
	}
}
