package ghapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/keitam913/textvid/domain"
)

type PostFileRepository struct {
	GitHubAPIPostRepositoryEndpoint string
}

func (pfRepo *PostFileRepository) ChangedFiles(event *domain.PushEvent) ([]*domain.PostFile, error) {
	res, err := http.Get(fmt.Sprintf("%s/compare/%s...%s", pfRepo.GitHubAPIPostRepositoryEndpoint, event.BeforeCommitID, event.AfterCommitID))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var comp CompareResponse
	if err := json.NewDecoder(res.Body).Decode(&comp); err != nil {
		return nil, err
	}

	var fs []*domain.PostFile
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

		fs = append(fs, &domain.PostFile{
			Filename: f.Filename,
			Content:  string(content),
		})
	}

	return fs, nil
}
