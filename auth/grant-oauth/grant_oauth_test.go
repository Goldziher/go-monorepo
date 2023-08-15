package grant_oauth_test

import (
	"context"
	"github.com/Goldziher/go-monorepo/db/mocks"
	"testing"

	"github.com/Goldziher/go-monorepo/auth/constants"
	grant_oauth "github.com/Goldziher/go-monorepo/auth/grant-oauth"
	"github.com/Goldziher/go-monorepo/db"
	"github.com/Goldziher/go-monorepo/lib/testutils"
	"github.com/stretchr/testify/assert"
)

func TestAuthInit(t *testing.T) {
	testutils.SetEnv(t)
	queryInstance := db.New(&mocks.QueryMock{})

	for _, testCase := range []struct {
		GrantType   string
		ExpectError bool
	}{
		{
			constants.GrantTypePassword,
			false,
		},
		{
			"code",
			true,
		},
	} {
		err := grant_oauth.AuthInit(context.TODO(), testCase.GrantType, queryInstance, db.UpsertUserParams{
			FullName:          testutils.TestingFullName,
			Email:             testutils.TestingEmail,
			PhoneNumber:       testutils.TestingPhoneNumber,
			ProfilePictureUrl: testutils.TestingUrl,
			Username:          testutils.TestingUsername,
			HashedPassword:    testutils.TestingPassword,
		})
		if testCase.ExpectError {
			assert.NotNil(t, err)
		} else {
			assert.Nil(t, err)
		}
	}
}

func TestGetUserData(t *testing.T) {
	testutils.SetEnv(t)
	queryInstance := db.New(&mocks.QueryMock{})

	for _, testCase := range []struct {
		GrantType   string
		ExpectError bool
	}{
		{
			constants.GrantTypePassword,
			false,
		},
		{
			"code",
			true,
		},
	} {
		user, err := grant_oauth.GetUserData(context.TODO(), testCase.GrantType, queryInstance, testutils.TestingEmail)
		if testCase.ExpectError {
			assert.NotNil(t, err)
			assert.Empty(t, user)
		} else {
			assert.Nil(t, err)
			assert.NotNil(t, user)
		}
	}
}
