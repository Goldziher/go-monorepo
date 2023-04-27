package api

import (
	"fmt"
	"net/http"

	"github.com/Goldziher/go-monorepo/auth/providers"
	"github.com/Goldziher/go-monorepo/auth/utils"
	"github.com/Goldziher/go-monorepo/lib/apiutils"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/rs/zerolog/log"
)

type OAuthInitRequestBody struct {
	Provider string `json:"provider"`
}

type OAuthInitResponseBody struct {
	RedirectUrl string `json:"redirectUrl"`
}

var (
	InitAuthPath = "/oauth/{provider}/init"
)

func InitOAuth(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")

	conf, providerErr := providers.GetProvider(provider)
	if providerErr != nil {
		log.Error().Err(providerErr).Str("provider", chi.URLParam(r, "provider")).Msg("unrecognized provider requested")
		_ = render.Render(w, r, apiutils.BadRequest(fmt.Sprintf("unsupported provider %v", provider)))
		return
	}

	state, stateErr := utils.CreateState()
	if stateErr != nil {
		log.Error().Err(stateErr).Msg("failed generating random state")
	}

	w.WriteHeader(200)
	render.JSON(w, r, OAuthInitResponseBody{RedirectUrl: conf.AuthCodeURL(state)})
}

func RegisterRoutes(router chi.Router) {
	router.Get(InitAuthPath, InitOAuth)
}
