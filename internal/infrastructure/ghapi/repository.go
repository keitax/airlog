package ghapi

import (
	"github.com/keitax/airlog/internal/domain"
)

type GitHubRepository struct{}

func (ghRepo *GitHubRepository) ChangedFiles(event *domain.PushEvent) ([]*domain.File, error) {
	return nil, nil
}
