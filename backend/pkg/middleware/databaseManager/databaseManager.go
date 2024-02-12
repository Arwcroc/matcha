package databaseManager

import (
	"github.com/gofiber/fiber/v2"
	"matcha/backend/pkg/database/arangodb"
	"matcha/backend/pkg/slog"
)

type Config struct {
	Filter func(c *fiber.Ctx) bool
	// Not using interface for Database here since it's hard to make it really generic (unlike SQL)
	Database *arangodb.DatabaseDriver
}

func NewHandler(config Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if c.Locals("database", config.Database) == nil {
			slog.Error("could not set database")
			return fiber.ErrInternalServerError
		}

		return c.Next()
	}
}
