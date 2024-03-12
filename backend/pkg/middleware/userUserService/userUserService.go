package userUserService

import (
	"github.com/gofiber/fiber/v2"
	"matcha/backend/pkg/database"
	"matcha/backend/pkg/object/user_user"
	"matcha/backend/pkg/slog"
)

func UserUserService(c *fiber.Ctx) error {
	driver := c.Locals("database").(database.Driver)

	userUserDriver, err := driver.NewObjectDriver(user_user.UserUser{})
	if err != nil {
		slog.Error(err)
		return fiber.ErrInternalServerError
	}
	if c.Locals("user_user_driver", userUserDriver) == nil {
		slog.Error("could not set user_user_driver")
		return fiber.ErrInternalServerError
	}

	return c.Next()
}
