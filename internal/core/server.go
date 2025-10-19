package core

import (
	"encoding/json"
	"net/http"
)

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
