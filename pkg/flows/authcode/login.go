package authcode

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"net/url"

	"github.com/deifyed/gatekeeper/pkg/core"
	"github.com/deifyed/gatekeeper/pkg/core/cookies"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type responseType string

const responseTypeCode responseType = "code"

type CreateLoginHandlerOpts struct {
	AuthorizeEndpoint string
	BaseURL           string
	ClientID          string
	CookieHandler     cookies.CookieHandler
	Logger            *logrus.Logger
}

func generateLoginRedirect(authorizeEndpoint, callbackURL, clientID, state string) (url.URL, error) {
	values := url.Values{}
	values.Add("response_type", string(responseTypeCode))
	values.Add("client_id", clientID)
	values.Add("redirect_uri", callbackURL)
	values.Add("scope", "offline_access openid")
	values.Add("state", state)

	redirect, err := url.Parse(authorizeEndpoint)
	if err != nil {
		return url.URL{}, fmt.Errorf("parsing authorization endpoint: %w", err)
	}

	redirect.RawQuery = values.Encode()

	return *redirect, nil
}

// CreateLoginHandler -
func CreateLoginHandler(storage core.StateStorage, opts CreateLoginHandlerOpts) gin.HandlerFunc {
	callbackURL := fmt.Sprintf("%s/callback", opts.BaseURL)

	logger := opts.Logger.WithFields(map[string]interface{}{
		"file": "flows/authcode/login.go",
		"func": "CreateLoginHandler",
	})

	return func(c *gin.Context) {
		logger.Debug("Starting login handler")

		stateID := uuid.New().String()
		state := uuid.New().String()

		logger.Debug(fmt.Sprintf("Storing \"%s\" with ID \"%s\"", state, stateID))

		err := storage.Put(stateID, state)
		if err != nil {
			logger.Error(fmt.Errorf("error storing state: %w", err))
			c.Status(http.StatusInternalServerError)

			return
		}

		redirect, err := generateLoginRedirect(opts.AuthorizeEndpoint, callbackURL, opts.ClientID, state)
		if err != nil {
			logger.Error(fmt.Errorf("error generating login redirect: %w", err))
			c.Status(http.StatusInternalServerError)

			return
		}

		opts.CookieHandler.SetStateID(c, stateID)

		c.Redirect(http.StatusFound, redirect.String())
		logger.Debug("Finished login handler")
	}
}
