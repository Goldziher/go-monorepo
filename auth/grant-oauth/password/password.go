package password

import (
	"context"
	"fmt"
	"github.com/Goldziher/go-monorepo/db"
	"github.com/Goldziher/go-monorepo/lib/database"
)

func CreateOrUpdateUser(ctx context.Context, user db.UpsertUserParams) error {
	dbConn := db.New(database.Get(ctx))
	_, err := dbConn.UpsertUser(ctx, user)

	if err != nil {
		return err
	}
	return nil
}

func GetUserData(ctx context.Context, email string) (db.User, error) {
	dbConn := db.New(database.Get(ctx))
	user, err := dbConn.GetUserByEmail(ctx, email)
	if err != nil {
		return db.User{}, fmt.Errorf(fmt.Sprintf("failed to fetch user details %+v", err))
	}
	return user, nil
}
