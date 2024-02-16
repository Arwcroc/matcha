package main

import (
	"github.com/arangodb/go-driver"
	"github.com/gofiber/fiber/v2"
	"matcha/backend/pkg/database/arangodb"
	"matcha/backend/pkg/middleware/databaseManager"
	"matcha/backend/pkg/middleware/logger"
	"matcha/backend/pkg/middleware/sessionManager"
	"matcha/backend/pkg/routes/auth"
	"matcha/backend/pkg/routes/photo"
	"matcha/backend/pkg/routes/user"
	"matcha/backend/pkg/slog"
	"matcha/backend/pkg/store/memory"
	"os"
	"os/signal"
	"syscall"
)

func startServer(app *fiber.App, bindAddress string) {
	err := app.Listen(bindAddress)
	slog.LogErrorExit(err)
}

func main() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	conf := fromEnv()
	slog.SetLogLevel(conf.loggerLevel)
	conf.print()

	memoryStore := memory.New()
	err := memoryStore.Connect()
	slog.LogErrorExit(err)
	defer memoryStore.Disconnect()
	sessions := sessionManager.New(memoryStore, sessionManager.Config{
		CookieKey: conf.cookieKey,
	})

	arango := arangodb.New(conf.dbConfig.url, "matcha",
		driver.BasicAuthentication(conf.dbConfig.username, conf.dbConfig.password),
	)
	slog.Info("Connecting to arangodb server")
	err = arango.Connect()
	slog.LogErrorExit(err)
	defer func() {
		slog.Info("Disconnecting from arangodb server")
		arango.Disconnect()
	}()
	_, err = arango.Database()
	slog.LogErrorExit(err)

	app := fiber.New()

	app.Use(databaseManager.NewHandler(databaseManager.Config{
		Database: &arango,
	}))
	app.Use(sessions.NewHandler())
	app.Use(logger.NewHandler(logger.Config{}))

	auth.Register(app)
	user.Register(app)
	photo.Register(app)

	go startServer(app, conf.bindAddress)
	slog.Info("Press Ctrl-c to shut down")
	<-c
	slog.Info("Ctrl-c pressed, shutting down...")
	_ = app.Shutdown()
}
