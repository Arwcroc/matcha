package photoService

import (
	"github.com/gofiber/fiber/v2"
	"matcha/backend/pkg/database"
	"matcha/backend/pkg/middleware/databaseManager"
	"matcha/backend/pkg/object"
	"matcha/backend/pkg/object/photo"
	"matcha/backend/pkg/slog"
)

const Local = "photo_object"

func PhotoService(c *fiber.Ctx) error {
	driver := c.Locals(databaseManager.Local).(database.Driver)

	photoObject, err := object.New[photo.Photo](driver)
	if err != nil {
		slog.Error(err)
		return fiber.ErrInternalServerError
	}
	if c.Locals(Local, photoObject) == nil {
		slog.Error("could not set " + Local)
		return fiber.ErrInternalServerError
	}

	return c.Next()
}
