package auth

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"matcha/backend/pkg/database"
	"matcha/backend/pkg/decorators"
	"matcha/backend/pkg/decorators/permissions"
	"matcha/backend/pkg/middleware/userService"
	"matcha/backend/pkg/object"
	"matcha/backend/pkg/slog"
	"matcha/backend/pkg/store"
	"matcha/backend/pkg/utils"
)

type credentials struct {
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password"`
}

func Register(app *fiber.App) {
	group := app.Group("/auth")
	group.Use(userService.UserService)
	group.Get("/whoami", decorators.Decorate(
		whoami,
		permissions.LoggedIn{},
	))
	group.Get("/logout", decorators.Decorate(
		logout,
		permissions.LoggedIn{},
	))
	group.Post("/login", login)
}

func login(c *fiber.Ctx) error {
	session := c.Locals("session").(*store.Session)
	userDriver := c.Locals("user_driver").(object.Driver)

	inputCredentials := credentials{}
	err := c.BodyParser(&inputCredentials)
	if err != nil {
		slog.Warn(err)
		return fiber.ErrBadRequest
	}

	key := "username"
	value := inputCredentials.Username
	if inputCredentials.Username == "" {
		key = "email"
		value = inputCredentials.Email
	}
	if key == "email" && inputCredentials.Email == "" {
		return fiber.ErrBadRequest
	}

	_, err = userDriver.Get(map[string]interface{}{
		key: value,
	})
	if err != nil {
		if errors.Is(err, database.NotFoundError) {
			return fiber.ErrNotFound
		}
		slog.Error(err)
		return fiber.ErrInternalServerError
	}

	if !utils.CheckPasswordHash(inputCredentials.Password, userDriver.GetField("password").(string)) {
		return fiber.ErrNotFound
	}

	userDriver.SetField("password", nil)
	session.Set("username", userDriver.GetField("username"))
	return c.JSON(*userDriver.GetInternal())
}

func logout(c *fiber.Ctx) error {
	session := c.Locals("session").(*store.Session)
	session.Delete("username")
	return c.SendStatus(fiber.StatusOK)
}

func whoami(c *fiber.Ctx) error {
	userDriver := c.Locals("user_driver").(object.Driver)
	session := c.Locals("session").(*store.Session)

	_, err := userDriver.Get(map[string]interface{}{
		"username": session.Get("username").(string),
	})
	if err != nil {
		if errors.Is(err, database.NotFoundError) {
			return fiber.ErrNotFound
		}
		slog.Error(err)
		return fiber.ErrInternalServerError
	}

	userDriver.SetField("password", nil)
	return c.JSON(*userDriver.GetInternal())
}
