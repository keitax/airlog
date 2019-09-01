package osenv

import (
	"fmt"
	"os"
	"strings"
)

type Config struct {
	Port                            string
	SiteTitle                       string
	Footnote                        string
	GitHubAPIPostRepositoryEndpoint string
}

func LoadConfig() (*Config, error) {
	conf := &Config{}
	var missed []string
	for _, prop := range []struct {
		Field *string
		Name  string
	}{
		{&conf.Port, "PORT"},
		{&conf.SiteTitle, "TV_SITE_TITLE"},
		{&conf.Footnote, "TV_FOOTNOTE"},
		{&conf.GitHubAPIPostRepositoryEndpoint, "TV_GITHUB_API_POST_REPOSITORY_ENDPOINT"},
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
