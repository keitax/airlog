//go:generate mockgen -package domain -source $GOFILE -destination mock_$GOFILE

package domain

import (
	"strings"
)

type PostService interface {
	GetByHTMLFilename(filename string) (*Post, error)
	Recent() ([]*Post, error)
}

type PostServiceImpl struct {
	Repository PostRepository
}

func (ps *PostServiceImpl) GetByHTMLFilename(filename string) (*Post, error) {
	markdownFilename := strings.Replace(filename, ".html", ".md", -1)
	p, err := ps.Repository.Filename(markdownFilename)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (ps *PostServiceImpl) Recent() ([]*Post, error) {
	recent, err := ps.Repository.All()
	if err != nil {
		return nil, err
	}
	return recent, nil
}
