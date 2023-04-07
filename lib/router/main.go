package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	chiMiddlewares "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog"
	"github.com/rs/zerolog/log"
)

func Create(serviceName string, middlewares ...func(http.Handler) http.Handler) chi.Router {
	router := chi.NewRouter()

	router.Use(httplog.RequestLogger(log.With().Str("service", serviceName).Logger()))
	router.Use(chiMiddlewares.Heartbeat("/health-check"))

	for _, middleware := range middlewares {
		router.Use(middleware)
	}

	return router
}
