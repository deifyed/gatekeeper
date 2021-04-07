package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"strings"
)

const tokenRefreshMiddlewarePath = "pkg/middleware/api.go"

func NewTokenRefreshMiddleware(opts NewTokenRefreshMiddlewareOpts) gin.HandlerFunc {
	logger := opts.Logger.WithFields(map[string]interface{}{
		"file": tokenRefreshMiddlewarePath,
		"func": "NewTokenRefreshMiddleware",
	})

	return func(c *gin.Context) { // TODO: algorithm needs work
		_, aerr := opts.CookieHandler.GetAccessToken(c)
		currentRefreshToken, rerr := opts.CookieHandler.GetRefreshToken(c)

		if aerr == nil || rerr != nil {
			logger.Debug("existing access token or missing refresh token. Nothing to do")

			c.Next()

			return
		}

		tr, err := requestToken(opts.TokenEndpoint, opts.ClientID, opts.ClientSecret, currentRefreshToken)
		if err != nil {
			logger.Warn(fmt.Errorf("error requesting new refresh token: %w", err))
		}

		opts.CookieHandler.SetAccessToken(c, tr.AccessToken, tr.ExpiresIn)
		opts.CookieHandler.SetRefreshToken(c, tr.RefreshToken)
		opts.CookieHandler.SetIDToken(c, tr.IDToken)
		opts.CookieHandler.SyncTokens(c)

		c.Next()
	}
}

func requestToken(tokenEndpoint, clientID, clientSecret, refreshToken string) (tokenResponse, error) {
	values := url.Values{}
	values.Add("grant_type", "refresh_token")
	values.Add("client_id", clientID)
	values.Add("client_secret", clientSecret)
	values.Add("refresh_token", refreshToken)

	response, err := http.Post(tokenEndpoint, "x-www-form-urlencoded", strings.NewReader(values.Encode()))
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
