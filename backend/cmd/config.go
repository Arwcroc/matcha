package main

import (
	"fmt"
	"matcha/backend/pkg/slog"
	"os"
)

type config struct {
	loggerLevel slog.Level
	bindAddress string
	cookieKey   string
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
		cookieKey:   envOrDefault("MATCHA__COOKIE_KEY", "session_key"),
	}
}

func (c *config) print() {
	slog.Info(fmt.Sprintf("MATCHA__LOG_LEVEL: 	%s", c.loggerLevel.String()))
	slog.Info(fmt.Sprintf("MATCHA__BIND_ADDRESS:	%s", c.bindAddress))
	slog.Info(fmt.Sprintf("MATCHA__COOKIE_KEY:	%s", c.cookieKey))
}
