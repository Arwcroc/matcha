package permissions

import (
	"github.com/gofiber/fiber/v2"
	"matcha/backend/pkg/decorators"
	"matcha/backend/pkg/store"
)

type Self struct {
	decorators.HandlerDecorator
	inner fiber.Handler
}

func (s Self) handler(c *fiber.Ctx) error {
	session := c.Locals("session").(*store.Session)
	sessionUsername := session.Get("username")
	paramUsername := c.Params("username")
	if sessionUsername != paramUsername {
		return fiber.ErrUnauthorized
	}
	return s.inner(c)
}

func (s Self) GetHandler() fiber.Handler {
	return LoggedIn{}.Decorate(s.handler).GetHandler()
}

func (s Self) Decorate(handler fiber.Handler) decorators.HandlerDecorator {
	s.inner = handler
	return s
}
