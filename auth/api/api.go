package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	grantOauth "github.com/Goldziher/go-monorepo/auth/grant-oauth"
	"github.com/Goldziher/go-monorepo/db"
	"github.com/Goldziher/go-monorepo/lib/cache"
	"github.com/Goldziher/go-monorepo/lib/hashing"

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

type GrantOAuthUserRequestBody struct {
	FullName          string `json:"full_name"`
	Email             string `json:"email"`
	PhoneNumber       string `json:"phone_number"`
	ProfilePictureUrl string `json:"profile_picture_url,omitempty"`
	Username          string `json:"username"`
	Password          string `json:"password"`
}

type GrantOAuthValidateRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GrantOAuthValidateResponseBody struct {
	FullName          string `json:"full_name"`
	Email             string `json:"email"`
	PhoneNumber       string `json:"phone_number"`
	ProfilePictureUrl string `json:"profile_picture_url"`
	Username          string `json:"username"`
}

const (
	InitAuthPath           = "/oauth/{provider}/init"
	OAuthCallbackPath      = "/oauth/{provider}/callback"
	OAuthGrantInitPath     = "/oauth/grant/{grant}/init"
	OAuthGrantValidatePath = "/oauth/grant/{grant}/validate"
)

func InitOAuth(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")

	conf, providerErr := providers.GetProvider(r.Context(), provider)
	if providerErr != nil {
		log.Error().
			Err(providerErr).
			Str("provider", chi.URLParam(r, "provider")).
			Msg("unrecognized provider requested")
		_ = render.Render(
			w,
			r,
			apiutils.BadRequest(fmt.Sprintf("unsupported provider %v", provider)),
		)
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
		log.Error().
			Err(providerErr).
			Str("provider", chi.URLParam(r, "provider")).
			Msg("unrecognized provider requested")
		_ = render.Render(
			w,
			r,
			apiutils.BadRequest(fmt.Sprintf("unsupported provider %v", provider)),
		)
		return
	}

	state, code := r.FormValue("state"), r.FormValue("code")

	log.Debug().
		Str("state", state).
		Str("code", code).
		Str("provider", provider).
		Msg("received oauth callback")

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

func InitGrantOAuth(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	grantType := chi.URLParam(r, "grant")

	var user GrantOAuthUserRequestBody
	if parseBodyErr := json.NewDecoder(r.Body).Decode(&user); parseBodyErr != nil {
		log.Error().Err(parseBodyErr).Msg("failed to parse the request payload")
		_ = render.Render(w, r, apiutils.BadRequest("invalid payload"))
		return
	}

	hashedPassword, hashingErr := hashing.Hash(user.Password)
	if hashingErr != nil {
		log.Error().Err(hashingErr).Msg("failed to encrypt the password longer than 72 bytes")
		_ = render.Render(w, r, apiutils.BadRequest("invalid password size"))
		return
	}

	authInitErr := grantOauth.AuthInit(ctx, grantType, db.GetQueries(), db.UpsertUserParams{
		FullName:          user.FullName,
		Email:             user.Email,
		PhoneNumber:       user.PhoneNumber,
		ProfilePictureUrl: user.ProfilePictureUrl,
		Username:          user.Username,
		HashedPassword:    hashedPassword,
	})
	if authInitErr != nil {
		log.Error().Err(authInitErr).Msg("failed to upsert the user details")
		_ = render.Render(w, r, apiutils.BadRequest("user signup failed"))
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, user.Email)
}

func GrantOAuthValidate(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	grantType := chi.URLParam(r, "grant")
	var credentials GrantOAuthValidateRequestBody
	if parseBodyErr := json.NewDecoder(r.Body).Decode(&credentials); parseBodyErr != nil {
		log.Error().Err(parseBodyErr).Msg("failed to parse the login credentials")
		_ = render.Render(w, r, apiutils.BadRequest("invalid payload"))
		return
	}

	user, authErr := grantOauth.GetUserData(ctx, grantType, db.GetQueries(), credentials.Email)
	if authErr != nil {
		log.Error().Err(authErr).Msg("user does not exist")
		_ = render.Render(w, r, apiutils.Unauthorized("user not found"))
		return
	}

	if isValid := hashing.CheckCode(credentials.Password, user.HashedPassword); !isValid {
		log.Error().Msg("incorrect password")
		_ = render.Render(w, r, apiutils.Unauthorized("incorrect credentials"))
		return
	}

	var userData GrantOAuthValidateResponseBody

	userResponseByte, _ := json.Marshal(user)
	if err := json.Unmarshal(userResponseByte, &userData); err != nil {
		log.Error().Err(err).Msg("error parsing the user data to response")
		_ = render.Render(w, r, apiutils.InternalServerError("failed to fetch the user data"))
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, userData)
}

func RegisterRoutes(router chi.Router) {
	router.Get(InitAuthPath, InitOAuth)
	router.Get(OAuthCallbackPath, OAuthCallback)
	router.Post(OAuthGrantInitPath, InitGrantOAuth)
	router.Post(OAuthGrantValidatePath, GrantOAuthValidate)
}
