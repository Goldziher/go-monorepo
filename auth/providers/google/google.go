package google

import (
	"context"
	"fmt"
	"sync"

	"github.com/Goldziher/go-monorepo/auth/config"
	"github.com/Goldziher/go-monorepo/auth/constants"
	"github.com/Goldziher/go-monorepo/auth/types"
	"github.com/Goldziher/go-monorepo/lib/apiutils"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/endpoints"
)

var (
	once   sync.Once
	google *oauth2.Config
)

const (
	UserProfileURL = "https://www.googleapis.com/oauth2/v1/userinfo?access_token=%s"
)

type GoogleUserData struct {
	Email         string `json:"email"`
	GivenName     string `json:"given_name"`
	Locale        string `json:"locale"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
	VerifiedEmail bool   `json:"verified_email"`
}

func GetConfig(ctx context.Context) (*oauth2.Config, error) {
	cfg, err := config.Get(ctx)
	if err != nil {
		return nil, err
	}

	once.Do(func() {
		google = &oauth2.Config{
			ClientID:     cfg.GoogleClientId,
			ClientSecret: cfg.GoogleClientSecret,
			RedirectURL:  fmt.Sprintf("%s/oauth/google/callback", cfg.BaseUrl),
			Scopes: []string{
				"https://www.googleapis.com/auth/userinfo.profile",
				"https://www.googleapis.com/auth/userinfo.email",
			},
			Endpoint: oauth2.Endpoint{
				AuthURL:  endpoints.Google.AuthURL,
				TokenURL: endpoints.Google.TokenURL,
			},
		}
	})
	return google, nil
}

func GetUserData(ctx context.Context, token *oauth2.Token) (*types.UserData, error) {
	client := google.Client(ctx, token)

	response, requestErr := client.Get(fmt.Sprintf(UserProfileURL, token.AccessToken))
	if requestErr != nil {
		return nil, requestErr
	}

	googleUserData := GoogleUserData{}
	deserializationError := apiutils.DeserializeJson(response, &googleUserData)
	if deserializationError != nil {
		return nil, deserializationError
	}

	return &types.UserData{
		Provider:          constants.ProviderGoogle,
		Email:             googleUserData.Email,
		FullName:          googleUserData.Name,
		Locale:            googleUserData.Locale,
		ProfilePictureUrl: googleUserData.Picture,
		Username:          googleUserData.GivenName,
		VerifiedEmail:     googleUserData.VerifiedEmail,
	}, nil
}
