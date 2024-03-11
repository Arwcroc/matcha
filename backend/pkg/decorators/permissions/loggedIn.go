package permissions

import (
	"github.com/gofiber/fiber/v2"
	"matcha/backend/pkg/decorators"
	"matcha/backend/pkg/store"
)

type LoggedIn struct {
	decorators.HandlerDecorator
	inner fiber.Handler
}

func (l LoggedIn) GetHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		session := c.Locals("session").(*store.Session)
		id := session.Get("username")
		if id == nil {
			return fiber.ErrUnauthorized
		}
		return l.inner(c)
	}
}

func (l LoggedIn) Decorate(handler fiber.Handler) decorators.HandlerDecorator {
	l.inner = handler
	return l
}
