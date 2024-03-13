package services

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"matcha/backend/pkg/database"
	"matcha/backend/pkg/decorators"
	"matcha/backend/pkg/middleware/sessionManager"
	"matcha/backend/pkg/middleware/userService"
	"matcha/backend/pkg/object/user"
	"matcha/backend/pkg/slog"
	"matcha/backend/pkg/store"
)

const SelfUserLocal = "self_user"

type SelfUser struct {
	decorators.HandlerDecorator
	inner fiber.Handler
}

func (s SelfUser) GetHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userObject := c.Locals(userService.Local).(user.User)
		session := c.Locals(sessionManager.Local).(*store.Session)

		o, err := userObject.Get(map[string]interface{}{
			"username": session.Get("username").(string),
		})
		if err != nil {
			if errors.Is(err, database.NotFoundError) {
				return fiber.ErrNotFound
			}
			slog.Error(err)
			return fiber.ErrInternalServerError
		}
		dbUser := o.(user.User)

		if c.Locals(SelfUserLocal, dbUser) == nil {
			slog.Error("could not set " + SelfUserLocal)
			return fiber.ErrInternalServerError
		}

		return s.inner(c)
	}
}

func (s SelfUser) Decorate(handler fiber.Handler) decorators.HandlerDecorator {
	s.inner = handler
	return s
}
