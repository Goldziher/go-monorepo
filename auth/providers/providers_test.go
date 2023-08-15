package providers_test

import (
	"context"
	"testing"

	"github.com/Goldziher/go-monorepo/lib/testutils"

	"github.com/Goldziher/go-monorepo/auth/constants"
	"github.com/Goldziher/go-monorepo/auth/providers"
	"github.com/stretchr/testify/assert"
)

func TestGetProvider(t *testing.T) {
	testutils.SetEnv(t)

	for _, testCase := range []struct {
		Provider    string
		ExpectError bool
	}{
		{
			constants.ProviderGithub,
			false,
		},
		{
			constants.ProviderGoogle,
			false,
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
