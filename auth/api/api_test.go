package api_test

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"testing"

	"github.com/Goldziher/go-monorepo/auth/api"
	"github.com/Goldziher/go-monorepo/lib/httpclient"
	"github.com/Goldziher/go-monorepo/lib/testutils"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

func createClient(t *testing.T) httpclient.Client {
	router := chi.NewRouter()
	api.RegisterRoutes(router)

	return testutils.CreateTestClient(t, router)
}

func TestInitOAuth(t *testing.T) {
	testutils.SetEnv(t)

	client := createClient(t)
	url := strings.ReplaceAll(api.InitAuthPath, "{provider}", "github")
	res, err := client.Get(context.TODO(), url)
	assert.Nil(t, err)
	assert.Equal(t, res.StatusCode, http.StatusOK)

	body := api.OAuthInitResponseBody{}

	unmarshalErr := json.Unmarshal(res.Body, &body)
	assert.Nil(t, unmarshalErr)
	assert.NotEmpty(t, body.RedirectUrl)
}
