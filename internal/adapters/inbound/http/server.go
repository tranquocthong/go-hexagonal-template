package http

import (
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"time"

	"example.com/yourorg/yourservice/internal/domain"
	portin "example.com/yourorg/yourservice/internal/domain/port/in"
	"example.com/yourorg/yourservice/pkg/config"
)

type Server struct {
	cfg      config.Config
	log      *slog.Logger
	greeting portin.GreetingUseCases
	http     *http.Server
}

func NewServer(cfg config.Config, log *slog.Logger, greeting portin.GreetingUseCases) *Server {
	mux := http.NewServeMux()
	s := &Server{cfg: cfg, log: log, greeting: greeting}

	mux.HandleFunc("GET /healthz", s.handleHealth)
	mux.HandleFunc("GET /api/v1/greetings", s.handleListGreetings)
	mux.HandleFunc("GET /api/v1/greetings/{id}", s.handleGetGreeting)
	mux.Handle("POST /api/v1/greetings", s.withAuth(http.HandlerFunc(s.handleCreateGreeting)))
	mux.HandleFunc("POST /api/v1/login", s.handleLogin)
	mux.Handle("GET /api/v1/me", s.withAuth(http.HandlerFunc(s.handleMe)))

	s.http = &http.Server{
		Addr:         cfg.HTTPAddress,
		Handler:      s.withLogging(s.withRecover(mux)),
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		IdleTimeout:  cfg.IdleTimeout,
	}
	return s
}

func (s *Server) Start() error { return s.http.ListenAndServe() }

func (s *Server) Stop(ctx context.Context) error { return s.http.Shutdown(ctx) }

func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	s.respondJSON(w, http.StatusOK, map[string]any{"status": "ok", "time": time.Now()})
}

func (s *Server) handleListGreetings(w http.ResponseWriter, r *http.Request) {
	items, err := s.greeting.ListGreetings()
	if err != nil {
		s.respondError(w, err)
		return
	}
	s.respondJSON(w, http.StatusOK, items)
}

func (s *Server) handleGetGreeting(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	g, err := s.greeting.GetGreeting(id)
	if err != nil {
		s.respondError(w, err)
		return
	}
	s.respondJSON(w, http.StatusOK, g)
}

type createGreetingRequest struct {
	ID      string `json:"id"`
	Message string `json:"message"`
}

func (s *Server) handleCreateGreeting(w http.ResponseWriter, r *http.Request) {
	var req createGreetingRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		s.respondJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid json"})
		return
	}
	g, err := s.greeting.CreateGreeting(req.ID, req.Message)
	if err != nil {
		s.respondError(w, err)
		return
	}
	s.respondJSON(w, http.StatusCreated, g)
}

func (s *Server) respondError(w http.ResponseWriter, err error) {
	code := http.StatusInternalServerError
	msg := "internal error"
	var de domain.DomainError
	if errors.As(err, &de) {
		switch de.Code {
		case domain.ErrNotFound:
			code = http.StatusNotFound
			msg = de.Message
		case domain.ErrAlreadyExists:
			code = http.StatusConflict
			msg = de.Message
		case domain.ErrInvalid:
			code = http.StatusBadRequest
			msg = de.Message
		case domain.ErrUnauthorized:
			code = http.StatusUnauthorized
			msg = de.Message
		case domain.ErrForbidden:
			code = http.StatusForbidden
			msg = de.Message
		default:
			code = http.StatusInternalServerError
			msg = de.Message
		}
	}
	s.respondJSON(w, code, map[string]string{"error": msg})
}

func (s *Server) respondJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func (s *Server) withLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rw := &responseRecorder{ResponseWriter: w, status: 200}
		next.ServeHTTP(rw, r)
		s.log.Info("http",
			slog.String("method", r.Method),
			slog.String("path", r.URL.Path),
			slog.Int("status", rw.status),
			slog.Duration("duration", time.Since(start)),
		)
	})
}

func (s *Server) withRecover(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				s.log.Error("panic recovered")
				s.respondJSON(w, http.StatusInternalServerError, map[string]string{"error": "internal error"})
			}
		}()
		next.ServeHTTP(w, r)
	})
}

type responseRecorder struct {
	http.ResponseWriter
	status int
}

func (rw *responseRecorder) WriteHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
}

var ErrServerClosed = http.ErrServerClosed
