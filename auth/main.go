package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Goldziher/go-monorepo/auth/config"
	"github.com/Goldziher/go-monorepo/db"

	"github.com/Goldziher/go-monorepo/auth/api"

	"github.com/Goldziher/go-monorepo/lib/logging"
	"github.com/Goldziher/go-monorepo/lib/router"
	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)

		<-c
		cancel()
	}()

	cfg, configParseErr := config.Get(ctx)

	if configParseErr != nil {
		log.Fatal().Err(configParseErr).Msg("failed to parse config, terminating")
	}

	dbConn := db.CreateConnection(ctx, cfg.DatabaseUrl)
	defer func() {
		_ = dbConn.Close(ctx)
	}()

	logging.Configure(cfg.Environment != "production")

	mux := router.Create("auth-service")

	api.RegisterRoutes(mux)

	httpServer := &http.Server{
		Addr:              fmt.Sprintf(":%d", cfg.Port),
		Handler:           mux,
		ReadHeaderTimeout: time.Second,
	}

	g, gCtx := errgroup.WithContext(ctx)

	g.Go(func() error {
		log.Info().Msg("server starting up")
		return httpServer.ListenAndServe()
	})

	g.Go(func() error {
		<-gCtx.Done()
		return httpServer.Shutdown(context.Background())
	})

	if err := g.Wait(); err != nil {
		log.Info().Msg(err.Error())
	}
}
