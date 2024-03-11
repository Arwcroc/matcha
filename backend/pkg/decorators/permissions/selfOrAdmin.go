package permissions

import (
	"github.com/gofiber/fiber/v2"
	"matcha/backend/pkg/decorators"
	"matcha/backend/pkg/store"
)

type SelfOrAdmin struct {
	decorators.HandlerDecorator
	inner fiber.Handler
}

func (s SelfOrAdmin) handler(c *fiber.Ctx) error {
	session := c.Locals("session").(*store.Session)
	sessionUsername := session.Get("username")
	role := session.Get("role")
	paramUsername := c.Params("username")
	if role != "admin" || sessionUsername != paramUsername {
		return fiber.ErrUnauthorized
	}
	return s.inner(c)
}

func (s SelfOrAdmin) GetHandler() fiber.Handler {
	return LoggedIn{}.Decorate(s.handler).GetHandler()
}

func (s SelfOrAdmin) Decorate(handler fiber.Handler) decorators.HandlerDecorator {
	s.inner = handler
	return s
}
