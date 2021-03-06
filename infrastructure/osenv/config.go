package osenv

import (
	"fmt"
	"os"
	"strings"
)

type Config struct {
	SiteTitle                       string
	Footnote                        string
	GitHubAPIPostRepositoryEndpoint string
	Mode                            string
}

func LoadConfig() (*Config, error) {
	conf := &Config{}
	var missed []string
	for _, prop := range []struct {
		Field *string
		Name  string
	}{
		{&conf.SiteTitle, "TV_SITE_TITLE"},
		{&conf.Footnote, "TV_FOOTNOTE"},
		{&conf.GitHubAPIPostRepositoryEndpoint, "TV_GITHUB_API_POST_REPOSITORY_ENDPOINT"},
		{&conf.Mode, "TV_MODE"},
	} {
		v := os.Getenv(prop.Name)
		if v == "" {
			missed = append(missed, prop.Name)
			continue
		}
		*prop.Field = v
	}
	if len(missed) > 0 {
		return nil, fmt.Errorf("missed env: %s", strings.Join(missed, ", "))
	}
	return conf, nil
}
