package http

import (
	"context"
	"net/http"
	"strings"

	"example.com/yourorg/yourservice/pkg/auth/jwt"
)

type authContextKey string

const ctxKeySubject authContextKey = "subject"

func (s *Server) withAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			s.respondJSON(w, http.StatusUnauthorized, map[string]string{"error": "missing bearer token"})
			return
		}
		token := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := jwt.Validate(token, s.cfg.JWTSecret)
		if err != nil {
			s.respondJSON(w, http.StatusUnauthorized, map[string]string{"error": "invalid token"})
			return
		}
		ctx := context.WithValue(r.Context(), ctxKeySubject, claims.Subject)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func subjectFromContext(ctx context.Context) (string, bool) {
	if v := ctx.Value(ctxKeySubject); v != nil {
		if s, ok := v.(string); ok {
			return s, true
		}
	}
	return "", false
}
