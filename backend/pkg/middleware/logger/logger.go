package logger

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"matcha/backend/pkg/slog"
	"time"
)

type Config struct {
	Filter func(c *fiber.Ctx) bool
}

func NewHandler(config Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if config.Filter != nil && config.Filter(c) {
			slog.Debug("Logger: Skipping middleware")
			return c.Next()
		}
		start := time.Now()
		next := c.Next()

		logString := fmt.Sprintf("%s -- %5s %s",
			c.IP(),
			c.Method(),
			c.Path(),
		)

		statusCode := c.Response().StatusCode()
		var fiberErr *fiber.Error
		if errors.As(next, &fiberErr) {
			statusCode = fiberErr.Code
		}

		logString = fmt.Sprintf("%s -> %d (%dμs, %dB)",
			logString,
			statusCode,
			time.Since(start).Microseconds(),
			len(c.Response().Body()),
		)

		slog.Info(logString)

		return next
	}
}
