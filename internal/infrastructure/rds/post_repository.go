package rds

import "github.com/keitax/airlog/internal/domain"

type PostRepository struct {
}

func (repo *PostRepository) Filename(filename string) (*domain.Post, error) {
	return &domain.Post{
		Filename: filename,
		Hash:     "xxx",
		Title:    "Title",
		Body: `# Title

## First post

Hello airlog!
`,
	}, nil
}

func (repo *PostRepository) All() ([]*domain.Post, error) {
	return nil, nil
}
