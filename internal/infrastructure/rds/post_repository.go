package rds

import "github.com/keitax/airlog/internal/domain"

type PostRepository struct {
}

func (repo *PostRepository) Filename(filename string) (*domain.Post, error) {
	return &domain.Post{
		Filename: filename,
		Hash:     "xxx",
		Title:    "Title",
		Body: `# Airlog

## h2

hello world

- a
- b
- c

### h3`,
	}, nil
}

func (repo *PostRepository) All() ([]*domain.Post, error) {
	return nil, nil
}
