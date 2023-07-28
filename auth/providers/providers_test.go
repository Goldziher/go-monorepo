package providers_test

import (
	"context"
	"github.com/Goldziher/go-monorepo/auth/providers"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetProvider(t *testing.T) {
	for _, testCase := range []struct {
		Provider    string
		ExpectError bool
	}{
		{
			providers.ProviderGithub,
			false,
		},
		{
			providers.ProviderGitlab,
			true,
		},
		{
			providers.ProviderBitBucket,
			true,
		},
		{
			providers.ProviderGoogle,
			true,
		},
		{
			"facebook",
			true,
		},
	} {
		config, err := providers.GetProvider(context.TODO(), testCase.Provider)
		if testCase.ExpectError {
			assert.Nil(t, config)
			assert.NotNil(t, err)
		} else {
			assert.NotNil(t, config)
			assert.Nil(t, err)
		}
	}
}
