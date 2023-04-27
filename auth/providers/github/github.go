package github

import (
	"fmt"
	"os"
	"sync"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/endpoints"
)

var (
	once   sync.Once
	github *oauth2.Config
)

func GetConfig() *oauth2.Config {
	once.Do(func() {
		github = &oauth2.Config{
			ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
			ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
			RedirectURL:  fmt.Sprintf("%s/oauth/github/redirect", os.Getenv("BASE_URL")),
			Scopes:       []string{"read:user", "user:email"},
			Endpoint: oauth2.Endpoint{
				AuthURL:  endpoints.GitHub.AuthURL,
				TokenURL: endpoints.GitHub.TokenURL,
			},
		}
	})
	return github
}
