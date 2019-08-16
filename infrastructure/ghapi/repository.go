package ghapi

import (
	"encoding/json"
	"fmt"
	"github.com/keitam913/airlog/domain"
	"io/ioutil"
	"net/http"
)

type GitHubRepository struct {
	GitHubAPIPostRepositoryEndpoint string
}

func (ghRepo *GitHubRepository) ChangedFiles(event *domain.PushEvent) ([]*domain.File, error) {
	res, err := http.Get(fmt.Sprintf("%s/compare/%s...%s", ghRepo.GitHubAPIPostRepositoryEndpoint, event.BeforeCommitID, event.AfterCommitID))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var comp CompareResponse
	if err := json.NewDecoder(res.Body).Decode(&comp); err != nil {
		return nil, err
	}

	var fs []*domain.File
	for _, f := range comp.Files {
		res, err := http.Get(f.RawURL)
		if err != nil {
			return nil, err
		}
		defer res.Body.Close()

		content, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		fs = append(fs, &domain.File{
			Path:    f.Filename,
			Content: string(content),
		})
	}

	return fs, nil
}
