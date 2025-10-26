package logger

import (
	"log/slog"
	"os"
	"strings"
)

func NewLogger(level, env string) *slog.Logger {
	var lvl slog.Level
	switch strings.ToLower(level) {
	case "debug":
		lvl = slog.LevelDebug
	case "warn":
		lvl = slog.LevelWarn
	case "error":
		lvl = slog.LevelError
	default:
		lvl = slog.LevelInfo
	}

	h := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: lvl})
	log := slog.New(h)
	return log.With(slog.String("env", env))
}
