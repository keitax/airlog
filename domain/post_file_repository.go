//go:generate mockgen -package domain -source $GOFILE -destination mock_$GOFILE

package domain

type PostFileRepository interface {
	ChangedFiles(event *PushEvent) ([]*PostFile, error)
}
