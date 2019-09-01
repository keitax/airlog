//go:generate mockgen -package $GOPACKAGE -source $GOFILE -destination mock_$GOFILE

package blog

import (
	"sort"
	"strings"

	"github.com/keitam913/textvid/domain"
)

type Service interface {
	GetByHTMLFilename(filename string) (*domain.Post, error)
	Recent() ([]*domain.Post, error)
	PushPosts(event *domain.PushEvent) error
}

type ServiceImpl struct {
	Service            domain.PostService
	Repository         domain.PostRepository
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
	sort.Slice(recent, func(i, j int) bool {
		return recent[i].Timestamp().After(recent[j].Timestamp())
	})
	return recent, nil
}

func (ps *ServiceImpl) PushPosts(event *domain.PushEvent) error {
	fs, err := ps.PostFileRepository.ChangedFiles(event)
	if err != nil {
		return err
	}
	for _, f := range fs {
		if domain.IsPostFileName(f.Filename) {
			post := ps.Service.ConvertToPost(f.Filename, f.Content)
			if err := ps.Repository.Put(post); err != nil {
				return err
			}
		}
	}
	return nil
}
