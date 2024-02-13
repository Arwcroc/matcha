package permissions

import (
	"github.com/gofiber/fiber/v2"
	"matcha/backend/pkg/slog"
	"matcha/backend/pkg/store"
)

func isSelf(c *fiber.Ctx) (bool, error) {
	session := c.Locals("session").(*store.Session)
	id := session.Get("username")
	username := c.Params("username")
	return id == username, nil
}

func Self(next fiber.Handler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		isSelf, err := isSelf(c)
		if err != nil {
			slog.Warn(err)
			return fiber.ErrBadRequest
		}

		if !isSelf {
			return fiber.ErrUnauthorized
		}
		return next(c)
	}
}
