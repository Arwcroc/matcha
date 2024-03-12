package decorators

import (
	"github.com/gofiber/fiber/v2"
)

type HandlerDecorator interface {
	GetHandler() fiber.Handler
	Decorate(handler fiber.Handler) HandlerDecorator
}

func Decorate(handler fiber.Handler, decorators ...HandlerDecorator) fiber.Handler {
	for _, decorator := range decorators {
		handler = decorator.Decorate(handler).GetHandler()
	}
	return handler
}
