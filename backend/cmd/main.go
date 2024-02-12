package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log/slog"
	"matcha/backend/pkg/middleware/logger"
	"matcha/backend/pkg/middleware/sessionManager"
	"matcha/backend/pkg/routes/auth"
	"matcha/backend/pkg/routes/user"
	"matcha/backend/pkg/store/memory"
	"os"
	"os/signal"
	"syscall"
)

func startServer(app *fiber.App, bindAddress string) {
	err := app.Listen(bindAddress)
	if err != nil {
		slog.Error(err.Error())
	}
}

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

func main() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	conf := fromEnv()
	slog.SetLogLoggerLevel(conf.loggerLevel)
	conf.print()

	memoryStore := memory.New()
	sessions := sessionManager.New(memoryStore, sessionManager.Config{
		CookieKey: conf.cookieKey,
	})

	app := fiber.New()

	auth.Register(app)
	user.Register(app)

	app.Use(logger.NewHandler(logger.Config{}))
	app.Use(sessions.NewHandler())

	go startServer(app, conf.bindAddress)
	slog.Info("Press Ctrl-c to shut down")
	<-c
	slog.Info("Ctrl-c pressed, shutting down...")
	_ = app.Shutdown()
}
