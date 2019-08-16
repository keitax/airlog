//go:generate mockgen -package domain -source $GOFILE -destination mock_$GOFILE

package domain

type PostRepository interface {
	Filename(filename string) (*Post, error)
	All() ([]*Post, error)
	Put(post *Post) error
}

type GitHubRepository interface {
	ChangedFiles(event *PushEvent) ([]*File, error)
}
