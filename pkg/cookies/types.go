package cookies

type CookieHandler struct {
	accessTokenKey  string
	refreshTokenKey string
	idTokenKey      string
	stateIDKey      string

	domain   string
	secure   bool
	httpOnly bool
}
