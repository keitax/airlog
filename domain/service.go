//go:generate mockgen -package domain -source $GOFILE -destination mock_$GOFILE

package domain

type PostService interface {
	ConvertToPost(filename, content string) *Post
}

type PostServiceImpl struct {
}

func (ps *PostServiceImpl) ConvertToPost(filename, content string) *Post {
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
	return post
}
