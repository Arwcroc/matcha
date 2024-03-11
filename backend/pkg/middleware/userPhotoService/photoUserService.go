package userPhotoService

import (
	"github.com/gofiber/fiber/v2"
	"matcha/backend/pkg/database"
	"matcha/backend/pkg/object/user_photo"
	"matcha/backend/pkg/slog"
)

func UserPhotoService(c *fiber.Ctx) error {
	driver := c.Locals("database").(database.Driver)

	userPhotoDriver, err := driver.NewObjectDriver(user_photo.UserPhoto{})
	if err != nil {
		slog.Error(err)
		return fiber.ErrInternalServerError
	}
	if c.Locals("user_photo_driver", userPhotoDriver) == nil {
		slog.Error("could not set user_photo_driver")
		return fiber.ErrInternalServerError
	}

	return c.Next()
}
