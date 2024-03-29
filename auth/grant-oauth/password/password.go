package password

import (
	"context"
	"fmt"

	"github.com/Goldziher/go-monorepo/db"
)

func CreateOrUpdateUser(ctx context.Context, q *db.Queries, user db.UpsertUserParams) error {
	_, err := q.UpsertUser(ctx, user)

	if err != nil {
		return err
	}
	return nil
}

func GetUserData(ctx context.Context, q *db.Queries, email string) (db.User, error) {
	user, err := q.GetUserByEmail(ctx, email)
	if err != nil {
		return db.User{}, fmt.Errorf(fmt.Sprintf("failed to fetch user details %+v", err))
	}
	return user, nil
}
