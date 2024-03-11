package userService

import (
	"github.com/gofiber/fiber/v2"
	"matcha/backend/pkg/database"
	"matcha/backend/pkg/object/user"
	"matcha/backend/pkg/slog"
)

func UserService(c *fiber.Ctx) error {
	driver := c.Locals("database").(database.Driver)

	userService, err := driver.NewObjectDriver(user.User{})
	if err != nil {
		return fiber.ErrInternalServerError
	}
	if c.Locals("user_driver", userService) == nil {
		slog.Error("could not set user_driver")
		return fiber.ErrInternalServerError
	}

	return c.Next()
}
