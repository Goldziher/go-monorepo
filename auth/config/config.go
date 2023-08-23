package config

import (
	"context"
	"sync"

	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	Port                  int    `env:"PORT,required"`
	Environment           string `env:"ENVIRONMENT,required"`
	BaseUrl               string `env:"BASE_URL,required"`
	GithubClientId        string `env:"GITHUB_CLIENT_ID,required"`
	GithubClientSecret    string `env:"GITHUB_CLIENT_SECRET,required"`
	GoogleClientId        string `env:"GOOGLE_CLIENT_ID,required"`
	GoogleClientSecret    string `env:"GOOGLE_CLIENT_SECRET,required"`
	MicrosoftClientId     string `env:"MICROSOFT_CLIENT_ID,required"`
	MicrosoftClientSecret string `env:"MICROSOFT_CLIENT_SECRET,required"`
	DatabaseUrl           string `env:"DATABASE_URL,required"`
}

var (
	config Config
	once   sync.Once
	err    error
)

func Get(ctx context.Context) (Config, error) {
	once.Do(func() {
		err = envconfig.Process(ctx, &config)
	})
	return config, err
}
