package handlers

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

const commonHandlersFile = "pkg/handlers/common.go"

func CreateLogoutHandler(opts CreateLogoutHandlerOpts) gin.HandlerFunc {
	logger := opts.Logger.WithFields(map[string]interface{}{
		"file": commonHandlersFile,
		"func": "CreateLogoutHandler",
	})

	return func(c *gin.Context) {
		values := url.Values{}

		refreshToken, err := opts.CookieHandler.GetRefreshToken(c.Request)
		if err != nil {
			logger.Warn(fmt.Errorf("unable to get refresh token: %w", err))

			c.Status(http.StatusBadRequest)

			return
		}

		values.Add("client_id", opts.ClientID)
		values.Add("client_secret", opts.ClientSecret)
		values.Add("refresh_token", refreshToken)

		response, err := http.Post(
			opts.LogoutEndpoint,
			"application/x-www-form-urlencoded",
			strings.NewReader(values.Encode()),
		)
		if err != nil {
			logger.Warn(fmt.Errorf("posting logout request: %w", err))

			c.Status(http.StatusInternalServerError)

			return
		}

		defer func() {
			_ = response.Body.Close()
		}()

		if response.StatusCode != http.StatusOK {
			logger.Warn(fmt.Errorf("bad response code logging out: %d", response.StatusCode))

			return
		}

		opts.CookieHandler.DeleteAccessToken(c.Writer)
		opts.CookieHandler.DeleteRefreshToken(c.Writer)
		opts.CookieHandler.DeleteIDToken(c.Writer)
	}
}
