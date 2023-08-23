package providers

import (
	"context"
	"fmt"

	"github.com/Goldziher/go-monorepo/auth/providers/microsoft"

	"github.com/Goldziher/go-monorepo/auth/constants"
	"github.com/Goldziher/go-monorepo/auth/providers/github"
	"github.com/Goldziher/go-monorepo/auth/providers/google"
	"github.com/Goldziher/go-monorepo/auth/types"
	"golang.org/x/oauth2"
)

func GetProvider(ctx context.Context, providerName string) (*oauth2.Config, error) {
	switch providerName {
	case constants.ProviderGithub:
		return github.GetConfig(ctx)
	case constants.ProviderGoogle:
		return google.GetConfig(ctx)
	case constants.ProviderMicrosoft:
		return microsoft.GetConfig(ctx)
	default:
		return nil, fmt.Errorf(fmt.Sprintf("unsupported provider %s", providerName))
	}
}

func GetUserData(ctx context.Context, token *oauth2.Token, providerName string) (*types.UserData, error) {
	switch providerName {
	case constants.ProviderGithub:
		return github.GetUserData(ctx, token)
	case constants.ProviderGoogle:
		return google.GetUserData(ctx, token)
	case constants.ProviderMicrosoft:
		return microsoft.GetUserData(ctx, token)
	default:
		return nil, fmt.Errorf(fmt.Sprintf("unsupported provider %s", providerName))
	}
}
