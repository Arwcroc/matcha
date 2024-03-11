package params

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"matcha/backend/pkg/database"
	"matcha/backend/pkg/decorators"
	"matcha/backend/pkg/object"
	"matcha/backend/pkg/slog"
)

type User struct {
	inner fiber.Handler
}

func (u User) GetHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userDriver := c.Locals("user_driver").(object.Driver)
		username := c.Params("username")

		_, err := userDriver.Get(map[string]interface{}{
			"username": username,
		})
		if err != nil {
			if errors.Is(err, database.NotFoundError) {
				return fiber.ErrNotFound
			}
			slog.Error(err)
			return fiber.ErrInternalServerError
		}

		if c.Locals("param_user", userDriver) == nil {
			slog.Error("could not set param_user")
			return fiber.ErrInternalServerError
		}

		return u.inner(c)
	}
}

func (u User) Decorate(handler fiber.Handler) decorators.HandlerDecorator {
	u.inner = handler
	return u
}
