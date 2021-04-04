package authcode

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/deifyed/gatekeeper/pkg/core"
	"github.com/deifyed/gatekeeper/pkg/core/cookies"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type grantType string

var grantTypeAuthorizationCode grantType = "authorization_code"

type CreateCallbackHandlerOpts struct {
	TokenEndpoint string
	ClientID      string
	ClientSecret  string
	CookieHandler cookies.CookieHandler
	Logger        *logrus.Logger
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
}

func CreateCallbackHandler(storage core.StateStorage, opts CreateCallbackHandlerOpts) gin.HandlerFunc {
	logger := opts.Logger.WithFields(map[string]interface{}{
		"file": "flows/authcode/callback.go",
		"func": "CreateCallbackHandler",
	})
	
	return func(c *gin.Context) {
		providedState := c.Query("state")
		providedCode := c.Query("code")

		providedStateID, err := opts.CookieHandler.GetStateID(c)
		if err != nil {
			logger.Warn(fmt.Errorf("error getting state ID from cookie: %w", err))
			
			c.Status(http.StatusBadRequest)

			return
		}

		err = validateState(storage, providedStateID, providedState)
		if err != nil {
			logger.Warn(
				fmt.Errorf("error validating state: %w", err),
				struct {
					providedID    string
					providedState string
					err           string
				}{
					providedID:    providedStateID,
					providedState: providedState,
					err:           err.Error(),
				})

			c.Status(http.StatusInternalServerError)
		}

		tokenResponse, err := postTokenRequest(opts.TokenEndpoint, providedCode, opts.ClientID, opts.ClientSecret)
		if err != nil {
			logger.Warn(fmt.Errorf("error posting token request: %w", err))

			c.Status(http.StatusInternalServerError)

			return
		}

		err = storage.Delete(providedStateID)
		if err != nil {
			logger.Warn(fmt.Errorf("error deleting stored state: %w", err))
		}

		opts.CookieHandler.DeleteStateID(c)

		opts.CookieHandler.SetAccessToken(c, tokenResponse.AccessToken, tokenResponse.ExpiresIn)
		opts.CookieHandler.SetRefreshToken(c, tokenResponse.RefreshToken)

		c.Redirect(http.StatusFound, "https://localhost:8000")
	}
}

func postTokenRequest(tokenEndpoint, code, clientID, clientSecret string) (TokenResponse, error) {
	values := url.Values{}
	values.Add("grant_type", string(grantTypeAuthorizationCode))
	values.Add("response_type", "token")
	values.Add("client_id", clientID)
	values.Add("client_secret", clientSecret)
	values.Add("code", code)

	request, err := http.NewRequest(http.MethodPost, tokenEndpoint, strings.NewReader(values.Encode()))
	if err != nil {
		return TokenResponse{}, fmt.Errorf("creating request: %w", err)
	}

	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return TokenResponse{}, fmt.Errorf("doing request: %w", err)
	}

	tokenResponse := TokenResponse{}

	err = json.NewDecoder(response.Body).Decode(&tokenResponse)
	if err != nil {
		return TokenResponse{}, fmt.Errorf("decoding json response: %w", err)
	}

	return tokenResponse, nil
}

func validateState(storage core.StateStorage, id, providedState string) error {
	state, err := storage.Get(id)
	if err != nil {
		return fmt.Errorf("getting state with id %s: %w", id, err)
	}

	if state != providedState {
		return fmt.Errorf("provided state did not pass state validation")
	}

	return nil
}
