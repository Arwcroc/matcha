package userService

import (
	"github.com/gofiber/fiber/v2"
	"matcha/backend/pkg/database"
	"matcha/backend/pkg/middleware/databaseManager"
	"matcha/backend/pkg/object"
	"matcha/backend/pkg/object/user"
	"matcha/backend/pkg/slog"
)

const Local = "user_object"

func UserService(c *fiber.Ctx) error {
	driver := c.Locals(databaseManager.Local).(database.Driver)

	userObject, err := object.New[user.User](driver)
	if err != nil {
		return fiber.ErrInternalServerError
	}
	if c.Locals(Local, userObject) == nil {
		slog.Error("could not set " + Local)
		return fiber.ErrInternalServerError
	}

	return c.Next()
}
