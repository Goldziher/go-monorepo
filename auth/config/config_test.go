package config_test

import (
	"context"
	"testing"

	"github.com/Goldziher/go-monorepo/auth/config"

	"github.com/stretchr/testify/assert"
)

func TestConfigGet(t *testing.T) {
	t.Run("successfully parses config", func(t *testing.T) {
		t.Setenv("PORT", "3000")
		t.Setenv("ENVIRONMENT", "development")
		t.Setenv("BASE_URL", "http://localhost")
		t.Setenv("GITHUB_CLIENT_ID", "githubClientId")
		t.Setenv("GITHUB_CLIENT_SECRET", "githubClientSecret")
		t.Setenv("GOOGLE_CLIENT_ID", "googleClientId")
		t.Setenv("GOOGLE_CLIENT_SECRET", "googleClientSecret")
		t.Setenv("DATABASE_URL", "databaseUrl")

		cfg, err := config.Get(context.TODO())
		assert.Nil(t, err)
		assert.Equal(t, cfg.Port, 3000)
		assert.Equal(t, cfg.Environment, "development")
		assert.Equal(t, cfg.BaseUrl, "http://localhost")
		assert.Equal(t, cfg.GithubClientId, "githubClientId")
		assert.Equal(t, cfg.GithubClientSecret, "githubClientSecret")
		assert.Equal(t, cfg.GoogleClientId, "googleClientId")
		assert.Equal(t, cfg.GoogleClientSecret, "googleClientSecret")
		assert.Equal(t, cfg.DatabaseUrl, "databaseUrl")
	})
}
