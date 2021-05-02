package handlers

import (
	"context"

	"github.com/coreos/go-oidc"
	"github.com/deifyed/gatekeeper/pkg/cookies"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
)

type CreateLoginHandlerOpts struct {
	CookieHandler cookies.CookieHandler
	Logger        *logrus.Logger
	Oauth2Config  oauth2.Config
}

type CreateCallbackHandlerOpts struct {
	Ctx           context.Context
	CookieHandler cookies.CookieHandler
	Logger        *logrus.Logger
	Oauth2Config  oauth2.Config
	TokenVerifier *oidc.IDTokenVerifier
}

type CreateUserinfoHandlerOpts struct {
	Ctx           context.Context
	CookieHandler cookies.CookieHandler
	Provider      *oidc.Provider
	Logger        *logrus.Logger
}

type CreateLogoutHandlerOpts struct {
	Logger         *logrus.Logger
	LogoutEndpoint string
	CookieHandler  cookies.CookieHandler

	ClientID     string
	ClientSecret string
}

type CreateProxyHandlerOpts struct {
	Logger *logrus.Logger

	CookieHandler cookies.CookieHandler

	Upstreams map[string]string
}

type tokenGetter struct {
	c             *gin.Context
	cookieHandler cookies.CookieHandler
}
