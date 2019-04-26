package osenv

import (
	"fmt"
	"github.com/keitax/airlog/internal/application"
	"os"
	"strings"
)

func LoadConfig() (*application.Config, error) {
	conf := &application.Config{}
	var missed []string
	for _, prop := range []struct {
		Field *string
		Name  string
	}{
		{&conf.Port, "PORT"},
		{&conf.SiteTitle, "AL_SITE_TITLE"},
		{&conf.Footnote, "AL_FOOTNOTE"},
		{&conf.BlogDSN, "AL_BLOG_DSN"},
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
