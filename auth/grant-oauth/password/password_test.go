package password_test

import (
	"context"
	"testing"

	"github.com/Goldziher/go-monorepo/db/mocks"

	"github.com/Goldziher/go-monorepo/lib/testutils"

	"github.com/Goldziher/go-monorepo/auth/grant-oauth/password"
	"github.com/Goldziher/go-monorepo/db"
	"github.com/stretchr/testify/assert"
)

func TestCreateOrUpdateUser(t *testing.T) {
	testutils.SetEnv(t)

	queryInstance := db.New(&mocks.QueryMock{})

	dbError := password.CreateOrUpdateUser(context.TODO(), queryInstance, db.UpsertUserParams{
		FullName:          testutils.TestingFullName,
		Email:             testutils.TestingEmail,
		PhoneNumber:       testutils.TestingPhoneNumber,
		ProfilePictureUrl: testutils.TestingUrl,
		Username:          testutils.TestingUsername,
		HashedPassword:    testutils.TestingPassword,
	})

	assert.Nil(t, dbError)
}

func TestGetUserData(t *testing.T) {
	t.Setenv("PORT", "3000")
	t.Setenv("ENVIRONMENT", "development")
	t.Setenv("BASE_URL", "http://localhost:3000")
	t.Setenv("DATABASE_URL", "postgresql://monorepo:monorepo@0.0.0.0:5432/monorepo?sslmode=disable")

	queryInstance := db.New(&mocks.QueryMock{})
	ctx := context.Background()

	user, err := password.GetUserData(ctx, queryInstance, testutils.TestingEmail)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, user.Email, testutils.TestingEmailResponse)
}
