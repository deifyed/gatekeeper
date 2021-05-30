package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

const tokenRefreshMiddlewarePath = "pkg/middleware/api.go" //nolint:gosec

type NextFunc func()

func NewTokenRefreshMiddleware(opts NewTokenRefreshMiddlewareOpts) http.HandlerFunc {
	logger := opts.Logger.WithFields(map[string]interface{}{
		"file": tokenRefreshMiddlewarePath,
		"func": "NewTokenRefreshMiddleware",
	})

	return func(w http.ResponseWriter, req *http.Request) { // TODO: algorithm needs work
		currentAccessToken, _ := opts.CookieHandler.GetAccessToken(req)
		currentRefreshToken, _ := opts.CookieHandler.GetRefreshToken(req)

		if currentAccessToken != "" || currentRefreshToken == "" {
			logger.Debug("existing access token or missing refresh token. Nothing to do")

			return
		}

		tr, err := requestToken(opts.TokenEndpoint, opts.ClientID, opts.ClientSecret, currentRefreshToken)
		if err != nil {
			logger.Warn(fmt.Errorf("error requesting new refresh token: %w", err))
		}

		opts.CookieHandler.SetAccessToken(w, tr.AccessToken, tr.ExpiresIn)
		opts.CookieHandler.SetRefreshToken(w, tr.RefreshToken)
		opts.CookieHandler.SetIDToken(w, tr.IDToken)
		opts.CookieHandler.SyncTokens(w, req)
	}
}

func requestToken(tokenEndpoint, clientID, clientSecret, refreshToken string) (tokenResponse, error) {
	values := url.Values{}
	values.Add("grant_type", "refresh_token")
	values.Add("client_id", clientID)
	values.Add("client_secret", clientSecret)
	values.Add("refresh_token", refreshToken)

	response, err := http.Post(tokenEndpoint, "x-www-form-urlencoded", strings.NewReader(values.Encode())) //nolint:gosec
	if err != nil {
		return tokenResponse{}, fmt.Errorf("requesting new refresh token: %w", err)
	}

	defer func() {
		_ = response.Body.Close()
	}()

	var tr tokenResponse

	err = json.NewDecoder(response.Body).Decode(&tr)
	if err != nil {
		return tokenResponse{}, fmt.Errorf("decoding json: %w", err)
	}

	return tr, nil
}
