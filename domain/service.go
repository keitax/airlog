//go:generate mockgen -package domain -source $GOFILE -destination mock_$GOFILE

package domain

import (
	"strings"
)

type PostService interface {
	GetByHTMLFilename(filename string) (*Post, error)
	Recent() ([]*Post, error)
	RegisterPost(filename, content string) error
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

func (ps *PostServiceImpl) RegisterPost(filename, content string) error {
	file := &PostFile{Filename: filename, Content: content}
	fm := file.ExtractFrontMatter()
	h1 := file.ExtractH1()
	post := &Post{
		Filename:  filename,
		Timestamp: file.GetTimestamp(),
		Title:     h1,
		Body:      file.Content,
	}
	if labels, ok := fm["labels"].([]interface{}); ok {
		for _, label := range labels {
			post.Labels = append(post.Labels, label.(string))
		}
	}
	return ps.Repository.Put(post)
}
