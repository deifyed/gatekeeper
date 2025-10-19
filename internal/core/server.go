package core

import (
	"encoding/json"
	"net/http"

	"github.com/deifyed/gatekeeper/gen"
)

type Server struct {
	router http.ServeMux
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *Server) respond(w http.ResponseWriter, _ *http.Request, data any, status int) {
	w.WriteHeader(status)

	if data != nil {
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)

			return
		}
	}
}

func (s *Server) decode(_ http.ResponseWriter, r *http.Request, v any) error {
	return json.NewDecoder(r.Body).Decode(v)
}

func (s *Server) GetLogin(w http.ResponseWriter, r *http.Request, params gen.GetLoginParams) {}

// Terminates a user session
// (POST /logout)
func (s *Server) PostLogout(w http.ResponseWriter, r *http.Request, params gen.PostLogoutParams) {}

// Acquire information about the user
// (GET /userinfo)
func (s *Server) GetUserinfo(w http.ResponseWriter, r *http.Request, params gen.GetUserinfoParams) {}
