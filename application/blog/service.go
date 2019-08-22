//go:generate mockgen -package $GOPACKAGE -source $GOFILE -destination mock_$GOFILE

package blog

import (
	"strings"

	"github.com/keitam913/airlog/domain"
)

type Service interface {
	GetByHTMLFilename(filename string) (*domain.Post, error)
	Recent() ([]*domain.Post, error)
	RegisterPost(filename, content string) error
	PushPost(event *domain.PushEvent) error
}

type ServiceImpl struct {
	Service          domain.PostService
	Repository       domain.PostRepository
	GitHubRepository domain.GitHubRepository
}

func (ps *ServiceImpl) GetByHTMLFilename(filename string) (*domain.Post, error) {
	markdownFilename := strings.Replace(filename, ".html", ".md", -1)
	p, err := ps.Repository.Filename(markdownFilename)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (ps *ServiceImpl) Recent() ([]*domain.Post, error) {
	recent, err := ps.Repository.All()
	if err != nil {
		return nil, err
	}
	return recent, nil
}

func (ps *ServiceImpl) RegisterPost(filename, content string) error {
	post := ps.Service.ConvertToPost(filename, content)
	return ps.Repository.Put(post)
}

func (ps *ServiceImpl) PushPost(event *domain.PushEvent) error {
	fs, err := ps.GitHubRepository.ChangedFiles(event)
	if err != nil {
		panic(err)
	}
	for _, f := range fs {
		if domain.IsPostFileName(f.Path) {
			post := ps.Service.ConvertToPost(f.Path, f.Content)
			if err := ps.Repository.Put(post); err != nil {
				return err
			}
		}
	}
	return nil
}
