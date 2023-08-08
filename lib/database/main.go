package database

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog/log"
	"os"
	"sync"
)

var (
	once   sync.Once
	dbConn *pgx.Conn
)

func Get(ctx context.Context) *pgx.Conn {
	once.Do(func() {
		conn, err := pgx.Connect(ctx, os.Getenv("DATABASE_URL"))
		if err != nil {
			log.Fatal().Err(err).Msg("failed to connect with database")
		}
		dbConn = conn
	})
	return dbConn
}

func Close(ctx context.Context) {
	err := dbConn.Close(ctx)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to disconnect the database")
	}
}
