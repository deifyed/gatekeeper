package middleware

import (
	"github.com/deifyed/gatekeeper/pkg/cookies"
	"github.com/sirupsen/logrus"
)

type NewTokenRefreshMiddlewareOpts struct {
	Logger        *logrus.Logger
	CookieHandler cookies.CookieHandler

	TokenEndpoint string
	ClientID      string
	ClientSecret  string
}

type tokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	IDToken      string `json:"id_token"`
	ExpiresIn    int    `json:"expires_in"`
}
