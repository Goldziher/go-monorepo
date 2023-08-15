package config_test

import (
	"context"
	"github.com/Goldziher/go-monorepo/lib/testutils"
	"testing"

	"github.com/Goldziher/go-monorepo/auth/config"

	"github.com/stretchr/testify/assert"
)

func TestConfigGet(t *testing.T) {
	t.Run("successfully parses config", func(t *testing.T) {
		testutils.SetEnv(t)

		cfg, err := config.Get(context.TODO())
		assert.Nil(t, err)
		assert.Equal(t, cfg.Port, 3000)
		assert.Equal(t, cfg.Environment, "development")
		assert.Equal(t, cfg.BaseUrl, "http://localhost")
		assert.Equal(t, cfg.GithubClientId, "githubClientId")
		assert.Equal(t, cfg.GithubClientSecret, "githubClientSecret")
		assert.Equal(t, cfg.GoogleClientId, "googleClientId")
		assert.Equal(t, cfg.GoogleClientSecret, "googleClientSecret")
		assert.Equal(t, cfg.DatabaseUrl, "postgresql://monorepo:monorepo@0.0.0.0:5432/monorepo?sslmode=disable")
	})
}
