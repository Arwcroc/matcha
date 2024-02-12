package permissions

import (
	"github.com/gofiber/fiber/v2"
	"matcha/backend/pkg/slog"
)

func SelfOrAdmin(next fiber.Handler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		isAdmin, _ := isAdmin(c)
		isSelf, err := isSelf(c)
		if err != nil {
			slog.Warn(err)
			return fiber.ErrBadRequest
		}

		if !isAdmin && !isSelf {
			return fiber.ErrUnauthorized
		}
		return next(c)
	}
}
