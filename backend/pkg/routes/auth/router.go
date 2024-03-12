package auth

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"matcha/backend/pkg/database"
	"matcha/backend/pkg/decorators"
	"matcha/backend/pkg/decorators/permissions"
	"matcha/backend/pkg/middleware/sessionManager"
	"matcha/backend/pkg/middleware/userService"
	"matcha/backend/pkg/object/user"
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
	session := c.Locals(sessionManager.Local).(*store.Session)
	userObject := c.Locals(userService.Local).(user.User)

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

	o, err := userObject.Get(map[string]interface{}{
		key: value,
	})
	if err != nil {
		if errors.Is(err, database.NotFoundError) {
			return fiber.ErrNotFound
		}
		slog.Error(err)
		return fiber.ErrInternalServerError
	}
	dbUser := o.(user.User)

	if !utils.CheckPasswordHash(inputCredentials.Password, dbUser.Password) {
		return fiber.ErrNotFound
	}

	dbUser.Password = ""
	session.Set("username", dbUser.Username)
	return c.JSON(dbUser)
}

func logout(c *fiber.Ctx) error {
	session := c.Locals(sessionManager.Local).(*store.Session)
	session.Delete("username")
	return c.SendStatus(fiber.StatusOK)
}

func whoami(c *fiber.Ctx) error {
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

	dbUser.Password = ""
	return c.JSON(dbUser)
}
