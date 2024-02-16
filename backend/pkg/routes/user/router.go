package user

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"matcha/backend/pkg/database"
	"matcha/backend/pkg/object"
	"matcha/backend/pkg/object/user"
	"matcha/backend/pkg/slog"
	"matcha/backend/pkg/utils"
)

// TODO find a way to put this in the database manager middleware (through decorators ?)
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

func getUserFromParam(next fiber.Handler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Accepts("json")

		userDriver := c.Locals("user_driver").(object.Driver)
		username := c.Params("username")

		_, err := userDriver.Get("username", username)
		if err != nil {
			if errors.Is(err, database.NotFoundError) {
				return fiber.ErrNotFound
			}
			slog.Error(err)
			return fiber.ErrInternalServerError
		}

		if c.Locals("param_user", userDriver) == nil {
			slog.Error("could not set param_user")
			return fiber.ErrInternalServerError
		}

		return next(c)
	}
}

func Register(app *fiber.App) {
	group := app.Group("/user")
	group.Use(getObjectDriver)
	group.Post("/", createUser)
	group.Get("/:username", getUserFromParam(getUser))
	group.Put("/:username", getUserFromParam(setUser))
	group.Delete("/:username", getUserFromParam(deleteUser))
}

func createUser(c *fiber.Ctx) error {
	c.Accepts("json")
	userDriver := c.Locals("user_driver").(object.Driver)

	inputUser := user.User{}
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

	err = userDriver.SetInternal(inputUser)
	if err != nil {
		slog.Warn(err)
		return fiber.ErrBadRequest
	}

	newUser, err := userDriver.Create()
	if err != nil {
		if errors.Is(err, database.UniqueConstraintError) {
			return fiber.ErrConflict
		}
		slog.Error(err)
		return fiber.ErrInternalServerError
	}

	userDriver.SetField("password", nil)
	return c.JSON(*newUser)
}

func getUser(c *fiber.Ctx) error {
	paramUser := c.Locals("param_user").(object.Driver)
	paramUser.SetField("password", nil)
	return c.JSON(*paramUser.GetInternal())
}

func setUser(c *fiber.Ctx) error {
	c.Accepts("json")
	paramUser := c.Locals("param_user").(object.Driver)

	inputUser := user.User{}
	err := c.BodyParser(&inputUser)
	if err != nil {
		slog.Warn(err)
		return fiber.ErrBadRequest
	}
	if !routes.CheckEmail(inputUser.Email) {
		return fiber.ErrBadRequest
	}
	if inputUser.Password == "" {
		inputUser.Password = paramUser.GetField("password").(string)
	}
	inputUser.Username = paramUser.GetField("username").(string)

	err = paramUser.SetInternal(inputUser)
	if err != nil {
		slog.Error(err)
		return fiber.ErrBadRequest
	}

	newUser, err := paramUser.Set()
	if err != nil {
		slog.Error(err)
		return fiber.ErrInternalServerError
	}

	paramUser.SetField("password", nil)
	return c.JSON(*newUser)
}

func deleteUser(c *fiber.Ctx) error {
	paramUser := c.Locals("param_user").(object.Driver)
	err := paramUser.Delete()
	if err != nil {
		slog.Error(err)
		return fiber.ErrInternalServerError
	}

	return c.SendStatus(fiber.StatusOK)
}
