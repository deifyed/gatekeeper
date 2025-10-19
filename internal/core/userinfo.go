package core

import (
	"net/http"

	"github.com/deifyed/gatekeeper/gen"
)

// GetUserinfo gets and returns userinfo related to a session
func (s *Server) GetUserinfo(w http.ResponseWriter, r *http.Request, params gen.GetUserinfoParams) {}
