package github_test

import (
	"context"
	"testing"

	"github.com/Goldziher/go-monorepo/lib/testutils"

	"github.com/Goldziher/go-monorepo/auth/providers/github"
	"github.com/stretchr/testify/assert"
)

func TestGetConfig(t *testing.T) {
	testutils.SetEnv(t)

	config, err := github.GetConfig(context.TODO())

	assert.Nil(t, err)
	assert.Equal(t, config.ClientID, "githubClientId")
	assert.Equal(t, config.ClientSecret, "githubClientSecret")
	assert.Equal(t, config.RedirectURL, "http://localhost/oauth/github/callback")
}
