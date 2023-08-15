package google_test

import (
	"context"
	"testing"

	"github.com/Goldziher/go-monorepo/lib/testutils"

	"github.com/Goldziher/go-monorepo/auth/providers/google"
	"github.com/stretchr/testify/assert"
)

func TestGetConfig(t *testing.T) {
	testutils.SetEnv(t)

	config, err := google.GetConfig(context.TODO())

	assert.Nil(t, err)
	assert.Equal(t, config.ClientID, "googleClientId")
	assert.Equal(t, config.ClientSecret, "googleClientSecret")
	assert.Equal(t, config.RedirectURL, "http://localhost/oauth/google/callback")
}
