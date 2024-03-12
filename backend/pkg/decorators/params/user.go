package params

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"matcha/backend/pkg/database"
	"matcha/backend/pkg/decorators"
	"matcha/backend/pkg/middleware/userService"
	"matcha/backend/pkg/object/user"
	"matcha/backend/pkg/slog"
)

const UserLocal = "param_user"

type User struct {
	inner fiber.Handler
}

func (u User) GetHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userObject := c.Locals(userService.Local).(user.User)
		username := c.Params("username")

		o, err := userObject.Get(map[string]interface{}{
			"username": username,
		})
		if err != nil {
			if errors.Is(err, database.NotFoundError) {
				return fiber.ErrNotFound
			}
			slog.Error(err)
			return fiber.ErrInternalServerError
		}

		if c.Locals(UserLocal, o.(user.User)) == nil {
			slog.Error("could not set " + UserLocal)
			return fiber.ErrInternalServerError
		}

		return u.inner(c)
	}
}

func (u User) Decorate(handler fiber.Handler) decorators.HandlerDecorator {
	u.inner = handler
	return u
}
