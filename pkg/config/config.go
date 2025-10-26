package config

import (
	"os"
	"strconv"
	"time"
)

type Config struct {
	AppName         string
	Env             string
	Version         string
	HTTPAddress     string
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	IdleTimeout     time.Duration
	ShutdownTimeout time.Duration
	LogLevel        string
	JWTSecret       string
	JWTTTL          time.Duration
}

func LoadFromEnv() Config {
	return Config{
		AppName:         getEnv("APP_NAME", "yourservice"),
		Env:             getEnv("APP_ENV", "development"),
		Version:         getEnv("APP_VERSION", "0.1.0"),
		HTTPAddress:     getEnv("HTTP_ADDRESS", ":8080"),
		ReadTimeout:     getEnvDuration("HTTP_READ_TIMEOUT", 15*time.Second),
		WriteTimeout:    getEnvDuration("HTTP_WRITE_TIMEOUT", 15*time.Second),
		IdleTimeout:     getEnvDuration("HTTP_IDLE_TIMEOUT", 60*time.Second),
		ShutdownTimeout: getEnvDuration("SHUTDOWN_TIMEOUT", 20*time.Second),
		LogLevel:        getEnv("LOG_LEVEL", "info"),
		JWTSecret:       getEnv("JWT_SECRET", "dev_secret_change_me"),
		JWTTTL:          getEnvDuration("JWT_TTL", time.Hour),
	}
}

func getEnv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

func getEnvDuration(key string, def time.Duration) time.Duration {
	if v := os.Getenv(key); v != "" {
		if n, err := strconv.Atoi(v); err == nil {
			return time.Duration(n) * time.Second
		}
		if d, err := time.ParseDuration(v); err == nil {
			return d
		}
	}
	return def
}
