package main

import (
	"github.com/gofiber/fiber/v2"
	"matcha/backend/pkg/middleware/logger"
	"matcha/backend/pkg/middleware/sessionManager"
	"matcha/backend/pkg/routes/auth"
	"matcha/backend/pkg/routes/user"
	"matcha/backend/pkg/slog"
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

func main() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	conf := fromEnv()
	slog.SetLogLevel(conf.loggerLevel)
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
