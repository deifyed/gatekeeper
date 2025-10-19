package core

import (
	"net/http"

	"github.com/deifyed/gatekeeper/gen"
)

// Terminates a user session
// (POST /logout)
func (s *Server) PostLogout(w http.ResponseWriter, r *http.Request, params gen.PostLogoutParams) {}
