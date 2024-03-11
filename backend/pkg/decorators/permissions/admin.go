package permissions

import (
	"github.com/gofiber/fiber/v2"
	"matcha/backend/pkg/decorators"
	"matcha/backend/pkg/store"
)

type Admin struct {
	decorators.HandlerDecorator
	inner fiber.Handler
}

func (a Admin) handler(c *fiber.Ctx) error {
	session := c.Locals("session").(*store.Session)
	role := session.Get("role")
	if role != "admin" {
		return fiber.ErrUnauthorized
	}
	return a.inner(c)
}

func (a Admin) GetHandler() fiber.Handler {
	return LoggedIn{}.Decorate(a.handler).GetHandler()
}

func (a Admin) Decorate(handler fiber.Handler) decorators.HandlerDecorator {
	a.inner = handler
	return a
}
