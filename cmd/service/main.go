package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"example.com/yourorg/yourservice/internal/adapters/inbound/http"
	outmem "example.com/yourorg/yourservice/internal/adapters/outbound/memory"
	"example.com/yourorg/yourservice/internal/app"
	"example.com/yourorg/yourservice/pkg/config"
	"example.com/yourorg/yourservice/pkg/logger"
)

// @title           Golang Template API
// @version         1.0
// @description     This is an API Server for the Golang Template application following DDD architecture.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @servers   http://localhost:8080

// @securityDefinitions.bearer BearerAuth
// @bearerFormat JWT
func main() {
	cfg := config.LoadFromEnv()
	log := logger.NewLogger(cfg.LogLevel, cfg.Env)
	slog.SetDefault(log)

	log.Info("starting service",
		slog.String("app", cfg.AppName),
		slog.String("env", cfg.Env),
		slog.String("version", cfg.Version),
		slog.String("addr", cfg.HTTPAddress),
	)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Wire dependencies
	repo := outmem.NewGreetingRepository()
	application := app.NewApplication(repo)

	server := http.NewServer(cfg, log, application)

	serverErrors := make(chan error, 1)
	go func() {
		serverErrors <- server.Start()
	}()

	select {
	case err := <-serverErrors:
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Error("server error", slog.String("error", err.Error()))
			os.Exit(1)
		}
	case <-ctx.Done():
		log.Info("shutdown signal received")
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), cfg.ShutdownTimeout)
	defer cancel()
	if err := server.Stop(shutdownCtx); err != nil {
		log.Error("graceful shutdown failed", slog.String("error", err.Error()))
		os.Exit(1)
	}

	log.Info("service stopped")
	fmt.Println("bye")
}
