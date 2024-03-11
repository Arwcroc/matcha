package decorators

import (
	"github.com/gofiber/fiber/v2"
)

type HandlerDecorator interface {
	GetHandler() fiber.Handler
	Decorate(handler fiber.Handler) HandlerDecorator
}
