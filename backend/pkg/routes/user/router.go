package user

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"matcha/backend/pkg/database"
	"matcha/backend/pkg/decorators"
	"matcha/backend/pkg/decorators/params"
	"matcha/backend/pkg/decorators/permissions"
	"matcha/backend/pkg/middleware/userService"
	"matcha/backend/pkg/object/user"
	"matcha/backend/pkg/routes"
	"matcha/backend/pkg/slog"
	"matcha/backend/pkg/utils"
)

func Register(app *fiber.App) {
	group := app.Group("/user")
	group.Use(userService.UserService)

	group.Post("/", createUser)
	group.Get("/:username", decorators.Decorate(
		getUser,
		permissions.LoggedIn{},
		params.User{},
	))
	group.Put("/:username", decorators.Decorate(
		setUser,
		permissions.SelfOrAdmin{},
		params.User{},
	))
	group.Delete("/:username", decorators.Decorate(
		deleteUser,
		permissions.SelfOrAdmin{},
		params.User{},
	))
}

func createUser(c *fiber.Ctx) error {
	c.Accepts("json")
	userObject := c.Locals(userService.Local).(user.User)

	inputUser := userObject
	err := c.BodyParser(&inputUser)
	if err != nil {
		slog.Warn(err)
		return fiber.ErrBadRequest
	}
	inputUser.Password, err = utils.HashPassword(inputUser.Password)
	if err != nil {
		slog.Error(err)
		return fiber.ErrInternalServerError
	}
	if !routes.CheckEmail(inputUser.Email) {
		return fiber.ErrBadRequest
	}

	o, err := inputUser.Create()
	if err != nil {
		if errors.Is(err, database.UniqueConstraintError) {
			return fiber.ErrConflict
		}
		slog.Error(err)
		return fiber.ErrInternalServerError
	}
	dbUser := o.(user.User)

	dbUser.Password = ""
	return c.JSON(dbUser)
}

func getUser(c *fiber.Ctx) error {
	paramUser := c.Locals(params.UserLocal).(user.User)
	paramUser.Password = ""
	return c.JSON(paramUser)
}

func setUser(c *fiber.Ctx) error {
	c.Accepts("json")
	userObject := c.Locals(userService.Local).(user.User)
	paramUser := c.Locals(params.UserLocal).(user.User)

	inputUser := userObject
	err := c.BodyParser(&inputUser)
	if err != nil {
		slog.Warn(err)
		return fiber.ErrBadRequest
	}
	if !routes.CheckEmail(inputUser.Email) {
		return fiber.ErrBadRequest
	}
	if inputUser.Password == "" {
		inputUser.Password = paramUser.Password
	}
	inputUser.Username = paramUser.Username

	o, err := inputUser.Set()
	if err != nil {
		slog.Error(err)
		return fiber.ErrInternalServerError
	}
	dbUser := o.(user.User)

	dbUser.Password = ""
	return c.JSON(dbUser)
}

func deleteUser(c *fiber.Ctx) error {
	paramUser := c.Locals(params.UserLocal).(user.User)
	err := paramUser.Delete()
	if err != nil {
		slog.Error(err)
		return fiber.ErrInternalServerError
	}

	return c.SendStatus(fiber.StatusOK)
}
