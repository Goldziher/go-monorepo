package testutils

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Goldziher/go-monorepo/lib/httpclient"
)

type TestClient struct {
}

func CreateTestClient(t *testing.T, handler http.Handler) httpclient.Client {
	server := httptest.NewServer(handler)

	t.Cleanup(func() {
		server.Close()
	})

	return httpclient.New(server.URL, server.Client())
}

func SetEnv(t *testing.T) {
	t.Setenv("PORT", "3000")
	t.Setenv("ENVIRONMENT", "development")
	t.Setenv("BASE_URL", "http://localhost")
	t.Setenv("GITHUB_CLIENT_ID", "githubClientId")
	t.Setenv("GITHUB_CLIENT_SECRET", "githubClientSecret")
	t.Setenv("GOOGLE_CLIENT_ID", "googleClientId")
	t.Setenv("GOOGLE_CLIENT_SECRET", "googleClientSecret")
	t.Setenv("DATABASE_URL", "postgresql://monorepo:monorepo@0.0.0.0:5432/monorepo?sslmode=disable")
	t.Setenv("MICROSOFT_CLIENT_ID", "microsoftClientId")
	t.Setenv("MICROSOFT_CLIENT_SECRET", "microsoftClientSecret")
}
