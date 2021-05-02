package middleware

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/deifyed/gatekeeper/pkg/cookies"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"golang.org/x/oauth2"
)

const (
	newAccessToken  = "refreshedAccessToken"
	newRefreshToken = "refreshedRefreshToken"
)

func TestNewTokenRefreshMiddleware(t *testing.T) {
	testCases := []struct {
		name string

		withTokens oauth2.Token

		expectTokens oauth2.Token
	}{
		{
			name: "Should do nothing with existing access token",

			withTokens: oauth2.Token{
				AccessToken:  "testtoken",
				RefreshToken: "",
			},

			expectTokens: oauth2.Token{
				AccessToken:  "",
				RefreshToken: "",
			},
		},
		{
			name: "Should do nothing with no access token nor refresh token",

			withTokens: oauth2.Token{},

			expectTokens: oauth2.Token{},
		},
		{
			name: "Should refresh tokens when refresh token is provided",

			withTokens: oauth2.Token{
				RefreshToken: "stillvalidtoken",
			},

			expectTokens: oauth2.Token{
				AccessToken:  newAccessToken,
				RefreshToken: newRefreshToken,
			},
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			cookieHandler := createCookieHandler()
			middlewareOpts := createMiddlewareOpts(cookieHandler)

			tokenServer := httptest.NewServer(mockTokenHandler{cookieHandler})
			middlewareOpts.TokenEndpoint = tokenServer.URL

			middleware := NewTokenRefreshMiddleware(middlewareOpts)

			recorder := httptest.NewRecorder()
			request := httptest.NewRequest(http.MethodGet, "http://example.com/", nil)

			request.AddCookie(&http.Cookie{
				Name:  "access_token",
				Value: tc.withTokens.AccessToken,
			})
			request.AddCookie(&http.Cookie{
				Name:  "refresh_token",
				Value: tc.withTokens.RefreshToken,
			})

			middleware(recorder, request)

			token := cookies.ExtractToken(recorder.Header().Values("Set-Cookie"))

			assert.Equal(t, tc.expectTokens.AccessToken, token.AccessToken)
			assert.Equal(t, tc.expectTokens.RefreshToken, token.RefreshToken)
		})
	}
}

func createCookieHandler() cookies.CookieHandler {
	return cookies.NewCookieHandler(
		"",
		"localhost",
		true,
		true,
	)
}

func createMiddlewareOpts(cookieHandler cookies.CookieHandler) NewTokenRefreshMiddlewareOpts {
	opts := NewTokenRefreshMiddlewareOpts{
		Logger:        logrus.New(),
		ClientID:      "dummyid",
		ClientSecret:  "dummysecret",
		CookieHandler: cookieHandler,
		TokenEndpoint: "",
	}
	opts.Logger.SetOutput(io.Discard)

	return opts
}

type mockTokenHandler struct {
	cookies.CookieHandler
}

func (m mockTokenHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	ctx, _ := gin.CreateTestContext(writer)
	ctx.Request = request

	token := oauth2.Token{
		AccessToken:  newAccessToken,
		RefreshToken: newRefreshToken,
	}

	payload, _ := json.Marshal(token)

	writer.Header().Set("Content-Type", "application/json")
	_, _ = writer.Write(payload)
}
