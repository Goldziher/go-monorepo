package config_test

import (
	"context"
	"testing"

	"github.com/Goldziher/go-monorepo/lib/config"
	"github.com/stretchr/testify/assert"
)

func TestConfigParse(t *testing.T) {
	t.Run("successfully parses config", func(t *testing.T) {
		t.Setenv("PORT", "3000")
		t.Setenv("ENVIRONMENT", "development")

		cfg, err := config.Parse(context.TODO())
		assert.Nil(t, err)
		assert.Equal(t, cfg.Port, 3000)
		assert.Equal(t, cfg.Environment, "development")
	})
}
