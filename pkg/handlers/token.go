package handlers

import (
	"fmt"
	"github.com/deifyed/gatekeeper/pkg/cookies"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"time"
)

func (t tokenGetter) Token() (*oauth2.Token, error) {
	accessToken, err := t.cookieHandler.GetAccessToken(t.c)
	if err != nil {
		return nil, fmt.Errorf("retrieving access token: %w", err)
	}

	refreshToken, err := t.cookieHandler.GetRefreshToken(t.c)
	if err != nil {
		return nil, fmt.Errorf("retrieving refresh token: %w", err)
	}

	return &oauth2.Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Expiry:       time.Now(),
	}, nil
}

func newTokenGetter(cookieHandler cookies.CookieHandler, c *gin.Context) oauth2.TokenSource {
	return &tokenGetter{
		cookieHandler: cookieHandler,
		c:             c,
	}
}
