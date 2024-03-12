package userUserService

import (
	"github.com/gofiber/fiber/v2"
	"matcha/backend/pkg/database"
	"matcha/backend/pkg/middleware/databaseManager"
	"matcha/backend/pkg/object"
	"matcha/backend/pkg/object/user_user"
	"matcha/backend/pkg/slog"
)

const Local = "user_user_object"

func UserUserService(c *fiber.Ctx) error {
	driver := c.Locals(databaseManager.Local).(database.Driver)

	userUserObject, err := object.New[user_user.UserUser](driver)
	if err != nil {
		slog.Error(err)
		return fiber.ErrInternalServerError
	}
	if c.Locals(Local, userUserObject) == nil {
		slog.Error("could not set " + Local)
		return fiber.ErrInternalServerError
	}

	return c.Next()
}
