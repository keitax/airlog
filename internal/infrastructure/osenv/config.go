package osenv

import (
	"github.com/keitax/airlog/internal/application"
	"os"
)

func LoadConfig() (*application.Config, error) {
	conf := &application.Config{}
	for _, prop := range []struct {
		Field *string
		Name  string
	}{
		{&conf.Port, "PORT"},
		{&conf.SiteTitle, "AL_SITE_TITLE"},
		{&conf.Footnote, "AL_FOOTNOTE"},
	} {
		*prop.Field = os.Getenv(prop.Name)
	}
	return conf, nil
}
