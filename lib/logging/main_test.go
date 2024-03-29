package logging_test

import (
	"testing"

	"github.com/Goldziher/go-monorepo/lib/logging"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestConfigLogger(t *testing.T) {
	t.Run("it sets global debug level when passed isDebug == true", func(t *testing.T) {
		logging.Configure(true)
		assert.Equal(t, zerolog.GlobalLevel(), zerolog.DebugLevel)
	})
	t.Run("it sets global info level when passed isDebug == false", func(t *testing.T) {
		logging.Configure(false)
		assert.Equal(t, zerolog.GlobalLevel(), zerolog.InfoLevel)
	})
}
