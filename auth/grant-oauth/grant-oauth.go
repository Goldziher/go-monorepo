package grant_oauth

import (
	"context"
	"fmt"

	"github.com/Goldziher/go-monorepo/auth/constants"
	"github.com/Goldziher/go-monorepo/auth/grant-oauth/password"
	"github.com/Goldziher/go-monorepo/db"
)

func AuthInit(ctx context.Context, grantType string, user db.UpsertUserParams) error {
	switch grantType {
	case constants.GrantTypePassword:
		return password.CreateOrUpdateUser(ctx, user)
	default:
		return fmt.Errorf(fmt.Sprintf("unsupported grant type %s", grantType))
	}
}

func GetUserData(ctx context.Context, grantType, email string) (db.User, error) {
	switch grantType {
	case constants.GrantTypePassword:
		return password.GetUserData(ctx, email)
	default:
		return db.User{}, fmt.Errorf(fmt.Sprintf("unsupported grant type %s", grantType))
	}
}
