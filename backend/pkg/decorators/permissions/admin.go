package permissions

import (
	"github.com/gofiber/fiber/v2"
	"matcha/backend/pkg/store"
)

func isAdmin(c *fiber.Ctx) (bool, error) {
	session := c.Locals("session").(*store.Session)
	role := session.Get("role")
	return role == "admin", nil
}

func Admin(next fiber.Handler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if isAdmin, _ := isAdmin(c); !isAdmin {
			return fiber.ErrUnauthorized
		}
		return next(c)
	}
}
