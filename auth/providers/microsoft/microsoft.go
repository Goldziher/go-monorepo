package microsoft

import (
	"context"
	"fmt"
	"sync"

	"github.com/Goldziher/go-monorepo/auth/config"
	"github.com/Goldziher/go-monorepo/auth/constants"
	"github.com/Goldziher/go-monorepo/auth/types"
	"github.com/Goldziher/go-monorepo/lib/apiutils"
	"github.com/rs/zerolog/log"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/endpoints"
)

var (
	once      sync.Once
	microsoft *oauth2.Config
)

type MicrosoftUserData struct {
	JobTitle          string `json:"jobTitle"`
	Email             string `json:"email"`
	DisplayName       string `json:"displayName"`
	PreferredLanguage string `json:"preferredLanguage"`
}

func GetConfig(ctx context.Context) (*oauth2.Config, error) {
	cfg, err := config.Get(ctx)
	if err != nil {
		return nil, err
	}

	once.Do(func() {
		microsoft = &oauth2.Config{
			ClientID:     cfg.MicrosoftClientId,
			ClientSecret: cfg.MicrosoftClientSecret,
			RedirectURL:  fmt.Sprintf("%s/oauth/microsoft/callback", cfg.BaseUrl),
			Scopes: []string{
				"User.Read",
			},
			Endpoint: oauth2.Endpoint{
				AuthURL:  endpoints.Microsoft.AuthURL,
				TokenURL: endpoints.Microsoft.TokenURL,
			},
		}
	})
	return microsoft, nil
}

func GetUserData(ctx context.Context, token *oauth2.Token) (*types.UserData, error) {
	client := microsoft.Client(ctx, token)

	response, requestErr := client.Get("https://graph.microsoft.com/v1.0/me")
	if requestErr != nil {
		return nil, requestErr
	}

	var microsoftUserData MicrosoftUserData

	deserializationError := apiutils.DeserializeJson(response, &microsoftUserData)
	if deserializationError != nil {
		return nil, deserializationError
	}
	log.Debug().
		Interface("microsoft user data", microsoftUserData).
		Msg("user data received from microsoft")

	return &types.UserData{
		Provider: constants.ProviderMicrosoft,
		Bio:      microsoftUserData.JobTitle,
		Email:    microsoftUserData.Email,
		FullName: microsoftUserData.DisplayName,
		Locale:   microsoftUserData.PreferredLanguage,
	}, nil
}
