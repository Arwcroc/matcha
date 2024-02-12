package logger

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log/slog"
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

		logString = fmt.Sprintf("%s -> %d (%dμs, %dB)",
			logString,
			c.Response().StatusCode(),
			time.Since(start).Microseconds(),
			len(c.Response().Body()),
		)

		println(logString)

		return next
	}
}
