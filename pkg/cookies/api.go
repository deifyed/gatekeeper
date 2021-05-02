package cookies

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"golang.org/x/oauth2"
)

const (
	monthMaxAge = int(time.Hour * 24 * 30)
)

func NewCookieHandler(cookiePrefix, domain string, secure, httpOnly bool) CookieHandler {
	return CookieHandler{
		accessTokenKey:  fmt.Sprintf("%saccess_token", cookiePrefix),
		refreshTokenKey: fmt.Sprintf("%srefresh_token", cookiePrefix),
		idTokenKey:      fmt.Sprintf("%sid_token", cookiePrefix),
		stateIDKey:      fmt.Sprintf("%sstate_id", cookiePrefix),
		domain:          domain,
		secure:          secure,
		httpOnly:        httpOnly,
	}
}

/*
 * Access token
 */
func (c CookieHandler) SetAccessToken(w http.ResponseWriter, accessToken string, expiresIn int) {
	c.setCookie(w, c.accessTokenKey, accessToken, "/", expiresIn)
}

func (c CookieHandler) GetAccessToken(request *http.Request) (string, error) {
	cookie, err := request.Cookie(c.accessTokenKey)
	if err != nil {
		return "", fmt.Errorf("getting access token cookie: %w", err)
	}

	return cookie.Value, nil
}

func (c CookieHandler) DeleteAccessToken(w http.ResponseWriter) {
	c.deleteCookie(w, c.accessTokenKey, "/")
}

/*
 * Refresh token
 */
func (c CookieHandler) SetRefreshToken(w http.ResponseWriter, refreshToken string) {
	c.setCookie(w, c.refreshTokenKey, refreshToken, "/", monthMaxAge)
}

func (c CookieHandler) GetRefreshToken(request *http.Request) (string, error) {
	cookie, err := request.Cookie(c.refreshTokenKey)
	if err != nil {
		return "", fmt.Errorf("getting refresh token cookie: %w", err)
	}

	return cookie.Value, nil
}

func (c CookieHandler) DeleteRefreshToken(w http.ResponseWriter) {
	c.deleteCookie(w, c.refreshTokenKey, "/")
}

/*
 * State ID
 */
func (c CookieHandler) SetStateID(w http.ResponseWriter, stateID string) {
	c.setCookie(w, c.stateIDKey, stateID, "/callback", int(time.Hour*1))
}

func (c CookieHandler) GetStateID(request *http.Request) (string, error) {
	cookie, err := request.Cookie(c.stateIDKey)
	if err != nil {
		return "", fmt.Errorf("getting state ID cookie: %w", err)
	}

	return cookie.Value, nil
}

func (c CookieHandler) DeleteStateID(w http.ResponseWriter) {
	c.deleteCookie(w, c.stateIDKey, "/callback")
}

/*
 * ID Token
 */
func (c CookieHandler) GetIDToken(request *http.Request) (string, error) {
	cookie, err := request.Cookie(c.idTokenKey)
	if err != nil {
		return "", fmt.Errorf("getting ID token cookie: %w", err)
	}

	return cookie.Value, nil
}

func (c CookieHandler) SetIDToken(w http.ResponseWriter, idToken string) {
	c.setCookie(w, c.idTokenKey, idToken, "/", monthMaxAge)
}

func (c CookieHandler) DeleteIDToken(w http.ResponseWriter) {
	c.deleteCookie(w, c.idTokenKey, "/")
}

// SyncTokens
func (c CookieHandler) SyncTokens(w http.ResponseWriter, req *http.Request) {
	token := ExtractToken(w.Header().Values("Set-Cookie"))

	maxAge := 3600 // TODO: Needs some smarter value maybe

	req.AddCookie(&http.Cookie{
		Name:     c.accessTokenKey,
		Value:    token.AccessToken,
		Path:     "/",
		Domain:   c.domain,
		MaxAge:   maxAge,
		Secure:   c.secure,
		HttpOnly: c.httpOnly,
	})

	req.AddCookie(&http.Cookie{
		Name:     c.refreshTokenKey,
		Value:    token.RefreshToken,
		Path:     "/",
		Domain:   c.domain,
		MaxAge:   maxAge,
		Secure:   c.secure,
		HttpOnly: c.httpOnly,
	})

	req.AddCookie(&http.Cookie{
		Name:     c.idTokenKey,
		Value:    "TODO",
		Path:     "/",
		Domain:   c.domain,
		MaxAge:   maxAge,
		Secure:   c.secure,
		HttpOnly: c.httpOnly,
	})
}

func (c CookieHandler) setCookie(w http.ResponseWriter, name, value, path string, expiresIn int) {
	http.SetCookie(w, &http.Cookie{
		Name:     name,
		Value:    value,
		Path:     path,
		Domain:   c.domain,
		MaxAge:   expiresIn,
		Secure:   c.secure,
		HttpOnly: c.httpOnly,
	})
}

func (c CookieHandler) deleteCookie(w http.ResponseWriter, name, path string) {
	http.SetCookie(w, &http.Cookie{
		Name:     name,
		Path:     path,
		Domain:   c.domain,
		MaxAge:   -1,
		Secure:   c.secure,
		HttpOnly: c.httpOnly,
	})
}

func ExtractToken(rawSetCookies []string) oauth2.Token {
	token := oauth2.Token{}

	if len(rawSetCookies) == 0 {
		return token
	}

	for _, cookie := range rawSetCookies {
		parts := strings.Split(cookie, ";")
		keyValueParts := strings.Split(parts[0], "=")

		key := keyValueParts[0]
		value := keyValueParts[1]

		switch key {
		case "access_token":
			token.AccessToken = value
		case "refresh_token":
			token.RefreshToken = value
		}
	}

	return token
}
