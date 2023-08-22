package microsoft_test

import (
	"context"
	"github.com/Goldziher/go-monorepo/auth/providers/microsoft"
	"github.com/Goldziher/go-monorepo/lib/testutils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetConfig(t *testing.T) {
	testutils.SetEnv(t)

	config, err := microsoft.GetConfig(context.TODO())

	assert.Nil(t, err)
	assert.Equal(t, config.ClientID, "microsoftClientId")
	assert.Equal(t, config.ClientSecret, "microsoftClientSecret")
	assert.Equal(t, config.RedirectURL, "http://localhost/oauth/microsoft/callback")
}
