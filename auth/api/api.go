package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
)

func Greet(w http.ResponseWriter, _ *http.Request) {
	if _, err := w.Write([]byte("hello world")); err != nil {
		log.Fatal().Err(err).Msg("failed to write message")
	}
}

func RegisterRoutes(router chi.Router) {
	router.Get("/", Greet)
}
