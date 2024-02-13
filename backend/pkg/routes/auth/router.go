package auth

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"matcha/backend/pkg/database"
	"matcha/backend/pkg/decorators/permissions"
	"matcha/backend/pkg/object"
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

func getObjectDriver(c *fiber.Ctx) error {
	driver := c.Locals("database").(database.Driver)

	userDriver, err := driver.NewObjectDriver(user.User{})
	if err != nil {
		return fiber.ErrInternalServerError
	}
	if c.Locals("user_driver", userDriver) == nil {
		slog.Error("could not set user_driver")
		return fiber.ErrInternalServerError
	}

	return c.Next()
}

func Register(app *fiber.App) {
	group := app.Group("/auth")
	group.Use(getObjectDriver)
	group.Get("/whoami", permissions.LoggedIn(whoami))
	group.Post("/login", login)
	group.Get("/logout", permissions.LoggedIn(logout))
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

	_, err = userDriver.Get(key, value)
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

	_, err := userDriver.Get("username", session.Get("username").(string))
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
