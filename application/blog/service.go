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
}

type ServiceImpl struct {
	Service    domain.PostService
	Repository domain.PostRepository
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
