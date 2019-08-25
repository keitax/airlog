//go:generate mockgen -package $GOPACKAGE -source $GOFILE -destination mock_$GOFILE

package blog

import (
	"strings"

	"github.com/keitam913/airlog/domain"
)

type Service interface {
	GetByHTMLFilename(filename string) (*domain.Post, error)
	Recent() ([]*domain.Post, error)
	PushPosts(event *domain.PushEvent) error
}

type ServiceImpl struct {
	Service          domain.PostService
	Repository       domain.PostRepository
	PostFileRepository domain.PostFileRepository
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

func (ps *ServiceImpl) PushPosts(event *domain.PushEvent) error {
	fs, err := ps.PostFileRepository.ChangedFiles(event)
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
