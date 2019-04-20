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
	} {
		*prop.Field = os.Getenv(prop.Name)
	}
	return conf, nil
}
