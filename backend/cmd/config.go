package main

import (
	"fmt"
	"matcha/backend/pkg/slog"
	"os"
)

type dbConfig struct {
	url      string
	username string
	password string
}

type sessionConfig struct {
	cookieKey string
}

type config struct {
	loggerLevel slog.Level
	bindAddress string
	sessionConfig
	dbConfig
}

func envOrDefault(key string, defaultValue string) string {
	value := defaultValue
	if envVal, ok := os.LookupEnv(key); ok {
		value = envVal
	}
	return value
}

func fromEnv() config {
	loggerLevel := slog.LevelInfo
	switch os.Getenv("MATCHA__LOG_LEVEL") {
	case slog.LevelDebug.String():
		loggerLevel = slog.LevelDebug
	case slog.LevelWarn.String():
		loggerLevel = slog.LevelWarn
	case slog.LevelError.String():
		loggerLevel = slog.LevelError
	}

	return config{
		loggerLevel: loggerLevel,
		bindAddress: envOrDefault("MATCHA__BIND_ADDRESS", "localhost:3000"),
		sessionConfig: sessionConfig{
			cookieKey: envOrDefault("MATCHA__SESSION__COOKIE_KEY", "session_key"),
		},
		dbConfig: dbConfig{
			url:      envOrDefault("MATCHA__DB__URL", "localhost:8529"),
			username: envOrDefault("MATCHA__DB__USERNAME", "root"),
			password: envOrDefault("MATCHA__DB__PASSWORD", "toor"),
		},
	}
}

func (c *config) print() {
	slog.Info(fmt.Sprintf("MATCHA__LOG_LEVEL: 	%s", c.loggerLevel.String()))
	slog.Info(fmt.Sprintf("MATCHA__BIND_ADDRESS:	%s", c.bindAddress))

	slog.Info(fmt.Sprintf("MATCHA__SESSION__COOKIE_KEY:	%s", c.cookieKey))

	slog.Info(fmt.Sprintf("MATCHA__DB__URL:	%s", c.dbConfig.url))
	slog.Info(fmt.Sprintf("MATCHA__DB__USERNAME:	%s", c.dbConfig.username))
	slog.Info(fmt.Sprintf("MATCHA__DB__PASSWORD:	%s", "<REDACTED>"))
}
