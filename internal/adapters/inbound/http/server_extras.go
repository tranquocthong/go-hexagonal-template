package http

import (
	"encoding/json"
	"io"
	"net/http"

	"example.com/yourorg/yourservice/pkg/auth/jwt"
)

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginResponse struct {
	AccessToken string `json:"access_token"`
}

func (s *Server) handleLogin(w http.ResponseWriter, r *http.Request) {
	// This is an example only; in real apps validate credentials properly.
	var req loginRequest
	if err := jsonNewDecoder(r.Body, &req); err != nil || req.Email == "" || req.Password == "" {
		s.respondJSON(w, http.StatusBadRequest, map[string]string{"error": "email and password required"})
		return
	}
	tok, err := jwt.Generate(req.Email, s.cfg.JWTSecret, s.cfg.JWTTTL)
	if err != nil {
		s.respondJSON(w, http.StatusInternalServerError, map[string]string{"error": "failed to create token"})
		return
	}
	s.respondJSON(w, http.StatusOK, loginResponse{AccessToken: tok})
}

func (s *Server) handleMe(w http.ResponseWriter, r *http.Request) {
	if sub, ok := subjectFromContext(r.Context()); ok {
		s.respondJSON(w, http.StatusOK, map[string]string{"subject": sub})
		return
	}
	s.respondJSON(w, http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
}

// tiny indirection for testing
func jsonNewDecoder(body io.ReadCloser, v any) error {
	defer body.Close()
	return json.NewDecoder(body).Decode(v)
}
