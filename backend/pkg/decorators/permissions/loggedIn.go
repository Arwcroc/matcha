package permissions

import (
	"github.com/gofiber/fiber/v2"
	"matcha/backend/pkg/store"
)

func isLoggedIn(c *fiber.Ctx) (bool, error) {
	session := c.Locals("session").(*store.Session)
	id := session.Get("username")
	return id != nil, nil
}

func LoggedIn(next fiber.Handler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		loggedIn, _ := isLoggedIn(c)
		if !loggedIn {
			return fiber.ErrUnauthorized
		}
		return next(c)
	}
}
