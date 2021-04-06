package cookies

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	deleteCookieAge = 0
	monthMaxAge     = int(time.Hour * 24 * 30)
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
func (c CookieHandler) SetAccessToken(ctx *gin.Context, accessToken string, expiresIn int) {
	ctx.SetCookie(c.accessTokenKey, accessToken, expiresIn, "/", c.domain, c.secure, c.httpOnly)
}

func (c CookieHandler) GetAccessToken(ctx *gin.Context) (string, error) {
	return ctx.Cookie(c.accessTokenKey)
}

func (c CookieHandler) DeleteAccessToken(ctx *gin.Context) {
	ctx.SetCookie(c.accessTokenKey, "", deleteCookieAge, "/", c.domain, c.secure, c.httpOnly)
}

/*
 * Refresh token
 */
func (c CookieHandler) SetRefreshToken(ctx *gin.Context, refreshToken string) {
	ctx.SetCookie(c.refreshTokenKey, refreshToken, monthMaxAge, "/", c.domain, c.secure, c.httpOnly)
}

func (c CookieHandler) GetRefreshToken(ctx *gin.Context) (string, error) {
	return ctx.Cookie(c.refreshTokenKey)
}

func (c CookieHandler) DeleteRefreshToken(ctx *gin.Context) {
	ctx.SetCookie(c.refreshTokenKey, "", deleteCookieAge, "/", c.domain, c.secure, c.httpOnly)
}

/*
 * State ID
 */
func (c CookieHandler) SetStateID(ctx *gin.Context, stateID string) {
	ctx.SetCookie(c.stateIDKey, stateID, int(time.Hour*1), "/callback", c.domain, c.secure, c.httpOnly)
}

func (c CookieHandler) GetStateID(ctx *gin.Context) (string, error) {
	return ctx.Cookie(c.stateIDKey)
}

func (c CookieHandler) DeleteStateID(ctx *gin.Context) {
	ctx.SetCookie(c.stateIDKey, "", deleteCookieAge, "/callback", c.domain, c.secure, c.httpOnly)
}

/*
 * ID Token
 */
func (c CookieHandler) GetIDToken(ctx *gin.Context) (string, error) {
	return ctx.Cookie(c.idTokenKey)
}

func (c CookieHandler) SetIDToken(ctx *gin.Context, idToken string) {
	ctx.SetCookie(c.idTokenKey, idToken, monthMaxAge, "/", c.domain, c.secure, c.httpOnly)
}

func (c CookieHandler) DeleteIDToken(ctx *gin.Context) {
	ctx.SetCookie(c.idTokenKey, "", deleteCookieAge, "/", c.domain, c.secure, c.httpOnly)
}
