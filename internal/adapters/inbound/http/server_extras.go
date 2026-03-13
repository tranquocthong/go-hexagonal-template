package http

import (
	"encoding/json"
	"io"
	"net/http"

	"example.com/yourorg/yourservice/pkg/auth/jwt"
)

type loginRequest struct {
	Email    string `json:"email" example:"user@example.com"`
	Password string `json:"password" example:"password123"`
}

type loginResponse struct {
	AccessToken string `json:"access_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI..."`
}

// handleLogin godoc
// @Summary      Login
// @Description  Login to get an access token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request body loginRequest true "Login Credentials"
// @Success      200 {object} loginResponse
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /api/v1/login [post]
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

// handleMe godoc
// @Summary      Get Profile
// @Description  Get current user profile
// @Tags         auth
// @Produce      json
// @Success      200 {object} map[string]string
// @Failure      401 {object} map[string]string
// @Security     BearerAuth
// @Router       /api/v1/me [get]
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
