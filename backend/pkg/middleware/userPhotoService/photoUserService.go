package userPhotoService

import (
	"github.com/gofiber/fiber/v2"
	"matcha/backend/pkg/database"
	"matcha/backend/pkg/middleware/databaseManager"
	"matcha/backend/pkg/object"
	"matcha/backend/pkg/object/user_photo"
	"matcha/backend/pkg/slog"
)

const Local = "user_photo_object"

func UserPhotoService(c *fiber.Ctx) error {
	driver := c.Locals(databaseManager.Local).(database.Driver)

	userPhotoObject, err := object.New[user_photo.UserPhoto](driver)
	if err != nil {
		slog.Error(err)
		return fiber.ErrInternalServerError
	}
	if c.Locals(Local, userPhotoObject) == nil {
		slog.Error("could not set " + Local)
		return fiber.ErrInternalServerError
	}

	return c.Next()
}
