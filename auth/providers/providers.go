package providers

import (
	"fmt"

	"github.com/Goldziher/go-monorepo/auth/providers/github"
	"golang.org/x/oauth2"
)

func GetProvider(providerName string) (*oauth2.Config, error) {
	switch providerName {
	case "github":
		return github.GetConfig(), nil
	case "gitlab":
		return nil, fmt.Errorf("not implemented")
	case "bitbucket":
		return nil, fmt.Errorf("not implemented")
	case "google":
		return nil, fmt.Errorf("not implemented")
	default:
		return nil, fmt.Errorf(fmt.Sprintf("unsupported provider %s", providerName))
	}
}
