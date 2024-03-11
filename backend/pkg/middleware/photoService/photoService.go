package photoService

import (
	"github.com/gofiber/fiber/v2"
	"matcha/backend/pkg/database"
	"matcha/backend/pkg/object/photo"
	"matcha/backend/pkg/slog"
)

func PhotoService(c *fiber.Ctx) error {
	driver := c.Locals("database").(database.Driver)

	photoDriver, err := driver.NewObjectDriver(photo.Photo{})
	if err != nil {
		slog.Error(err)
		return fiber.ErrInternalServerError
	}
	if c.Locals("photo_driver", photoDriver) == nil {
		slog.Error("could not set photo_driver")
		return fiber.ErrInternalServerError
	}

	return c.Next()
}
