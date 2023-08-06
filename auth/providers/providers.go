package providers

import (
	"context"
	"fmt"
	"github.com/Goldziher/go-monorepo/auth/constants"
	"github.com/Goldziher/go-monorepo/auth/providers/github"
	"github.com/Goldziher/go-monorepo/auth/types"
	"golang.org/x/oauth2"
)

func GetProvider(ctx context.Context, providerName string) (*oauth2.Config, error) {
	switch providerName {
	case constants.ProviderGithub:
		return github.GetConfig(ctx)
	case constants.ProviderGitlab:
	case constants.ProviderBitBucket:
	case constants.ProviderGoogle:
		return nil, fmt.Errorf(fmt.Sprintf("not implemented for provider %s", providerName))
	}
	return nil, fmt.Errorf(fmt.Sprintf("unsupported provider %s", providerName))
}

func GetUserData(ctx context.Context, token *oauth2.Token, providerName string) (*types.UserData, error) {
	switch providerName {
	case constants.ProviderGithub:
		return github.GetUserData(ctx, token)
	case constants.ProviderGitlab:
	case constants.ProviderBitBucket:
	case constants.ProviderGoogle:
		return nil, fmt.Errorf(fmt.Sprintf("not implemented for provider %s", providerName))
	}
	return nil, fmt.Errorf(fmt.Sprintf("unsupported provider %s", providerName))
}
