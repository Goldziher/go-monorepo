package api

import (
	"fmt"
	"github.com/Goldziher/go-monorepo/lib/cache"
	"net/http"
	"time"

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

const (
	InitAuthPath      = "/oauth/{provider}/init"
	OAuthCallbackPath = "/oauth/{provider}/callback"
)

func InitOAuth(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")

	conf, providerErr := providers.GetProvider(r.Context(), provider)
	if providerErr != nil {
		log.Error().Err(providerErr).Str("provider", chi.URLParam(r, "provider")).Msg("unrecognized provider requested")
		_ = render.Render(w, r, apiutils.BadRequest(fmt.Sprintf("unsupported provider %v", provider)))
		return
	}
	log.Info().Str("redirect-url", conf.RedirectURL).Msg("using redirect-url")
	state := utils.CreateStateString()

	cacheErr := cache.Get().Set(r.Context(), state, state, 10*time.Second).Err()

	if cacheErr != nil {
		log.Error().Err(cacheErr).Msg("failed to cache data in redis")
		_ = render.Render(w, r, apiutils.InternalServerError("failed to cache state"))
		return
	}

	http.Redirect(w, r, conf.AuthCodeURL(state), http.StatusTemporaryRedirect)
}

func OAuthCallback(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")
	conf, providerErr := providers.GetProvider(r.Context(), provider)
	if providerErr != nil {
		log.Error().Err(providerErr).Str("provider", chi.URLParam(r, "provider")).Msg("unrecognized provider requested")
		_ = render.Render(w, r, apiutils.BadRequest(fmt.Sprintf("unsupported provider %v", provider)))
		return
	}

	state, code := r.FormValue("state"), r.FormValue("code")

	log.Debug().Str("state", state).Str("code", code).Str("provider", provider).Msg("received oauth callback")

	defer func() {
		_ = cache.Get().Del(r.Context(), state)
	}()

	cacheErr := cache.Get().Get(r.Context(), state).Err()
	if cacheErr != nil {
		log.Error().Err(cacheErr).Str("state", state).Msg("failed to retrieve cached data")
		_ = render.Render(w, r, apiutils.Unauthorized("state validation failed"))
		return
	}

	token, err := conf.Exchange(r.Context(), code)
	if err != nil {
		log.Error().Err(err).Msg("failed to retrieve auth token from provider")
		_ = render.Render(w, r, apiutils.Unauthorized("token retrieval failed"))
	}

	if !token.Valid() {
		log.Error().Str("token", token.AccessToken).Msg("invalid token")
		_ = render.Render(w, r, apiutils.Unauthorized("token is invalid"))
	}

	userData, getUserErr := providers.GetUserData(r.Context(), token, provider)
	if getUserErr != nil {
		log.Error().Err(getUserErr).Msg("failed to retrieve user from provider")
		_ = render.Render(w, r, apiutils.Unauthorized("user retrieval failed"))
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, userData)
}

func RegisterRoutes(router chi.Router) {
	router.Get(InitAuthPath, InitOAuth)
	router.Get(OAuthCallbackPath, OAuthCallback)
}
