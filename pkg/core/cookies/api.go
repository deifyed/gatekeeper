package cookies

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func NewCookieHandler(cookiePrefix, domain string, secure, httpOnly bool) CookieHandler {
	return CookieHandler{
		accessTokenKey:  fmt.Sprintf("%saccess_token", cookiePrefix),
		refreshTokenKey: fmt.Sprintf("%srefresh_token", cookiePrefix),
		stateIDKey:      fmt.Sprintf("%sstate_id", cookiePrefix),
		domain:          domain,
		secure:          secure,
		httpOnly:        httpOnly,
	}
}

func (c CookieHandler) SetAccessToken(ctx *gin.Context, accessToken string, expiresIn int) {
	ctx.SetCookie(
		c.accessTokenKey,
		accessToken,
		expiresIn,
		"/",
		c.domain,
		c.secure,
		c.httpOnly,
	)
}

func (c CookieHandler) GetAccessToken(ctx *gin.Context) (string, error) {
	return ctx.Cookie(c.accessTokenKey)
}

func (c CookieHandler) SetRefreshToken(ctx *gin.Context, refreshToken string) {
	ctx.SetCookie(
		c.refreshTokenKey,
		refreshToken,
		0,
		"/",
		c.domain,
		c.secure,
		c.httpOnly,
	)
}

func (c CookieHandler) GetRefreshToken(ctx *gin.Context) (string, error) {
	return ctx.Cookie(c.refreshTokenKey)
}

func (c CookieHandler) SetStateID(ctx *gin.Context, stateID string) {
	ctx.SetCookie(
		c.stateIDKey,
		stateID,
		int(time.Hour*1),
		"/callback",
		c.domain,
		c.secure,
		c.httpOnly,
	)
}

func (c CookieHandler) GetStateID(ctx *gin.Context) (string, error) {
	return ctx.Cookie(c.stateIDKey)
}

func (c CookieHandler) DeleteStateID(ctx *gin.Context) {
	ctx.SetCookie(c.stateIDKey, "", -1, "/callback", c.domain, c.secure, c.httpOnly)
}
