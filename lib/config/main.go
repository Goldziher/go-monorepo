package config

import (
	"context"
	"sync"

	envconfig "github.com/sethvargo/go-envconfig"
)

type Config struct {
	Port  int  `env:"PORT"`
	Debug bool `env:"DEBUG,default=false"`
}

var (
	config Config
	once   sync.Once
	err    error
)

func Parse(ctx context.Context) (Config, error) {
	once.Do(func() {
		err = envconfig.Process(ctx, &config)
	})
	return config, err
}
