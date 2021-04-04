package cookies

type CookieHandler struct {
	accessTokenKey  string
	refreshTokenKey string
	stateIDKey      string

	domain   string
	secure   bool
	httpOnly bool
}
