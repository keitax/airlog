package dynamodb

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/keitam913/textvid/domain"
)

const PostTableName = "Post"

type PostRepository struct {
	DB *dynamodb.DynamoDB
}

func (pr *PostRepository) Filename(filename string) (*domain.Post, error) {
	out, err := pr.DB.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(PostTableName),
		Key: map[string]*dynamodb.AttributeValue{
			"Filename": {
				S: aws.String(filename),
			},
		},
	})
	if err != nil {
		return nil, err
	}
	return itemToPost(out.Item), nil
}

func (pr *PostRepository) All() ([]*domain.Post, error) {
	out, err := pr.DB.Scan(&dynamodb.ScanInput{
		TableName: aws.String(PostTableName),
	})
	if err != nil {
		return nil, err
	}
	var posts []*domain.Post
	for _, item := range out.Items {
		posts = append(posts, itemToPost(item))
	}
	return posts, nil
}

func (pr *PostRepository) Put(post *domain.Post) error {
	labels := []*dynamodb.AttributeValue{}
	for _, l := range post.Labels {
		labels = append(labels, &dynamodb.AttributeValue{S: aws.String(l)})
	}
	if _, err := pr.DB.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(PostTableName),
		Item: map[string]*dynamodb.AttributeValue{
			"Filename": {S: aws.String(post.Filename)},
			"Title":    {S: aws.String(post.Title)},
			"Body":     {S: aws.String(post.Body)},
			"Labels":   {L: labels},
		},
	}); err != nil {
		return err
	}
	return nil
}

func itemToPost(item map[string]*dynamodb.AttributeValue) *domain.Post {
	p := &domain.Post{
		Filename: *item["Filename"].S,
		Title:    *item["Title"].S,
		Body:     *item["Body"].S,
	}
	for _, l := range item["Labels"].L {
		p.Labels = append(p.Labels, *l.S)
	}
	return p
}
