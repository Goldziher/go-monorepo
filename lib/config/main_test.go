package config_test

import (
	"context"
	"os"
	"testing"

	"github.com/Goldziher/go-monorepo/lib/config"
	"github.com/stretchr/testify/assert"
)

func TestConfigParse(t *testing.T) {
	t.Run("successfully parses config", func(t *testing.T) {
		_ = os.Setenv("PORT", "3000")

		cfg, err := config.Parse(context.TODO())
		assert.Nil(t, err)
		assert.Equal(t, cfg.Port, 3000)
	})
}
