package user

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"matcha/backend/pkg/database"
	"matcha/backend/pkg/object"
	"matcha/backend/pkg/object/user"
	"matcha/backend/pkg/slog"
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

func Register(app *fiber.App) {
	group := app.Group("/user")
	group.Use(getObjectDriver)
	group.Post("/", createUser)
	group.Get("/:id", getUser)
	group.Put("/:id", setUser)
	group.Delete("/:id", deleteUser)
}

func createUser(c *fiber.Ctx) error {
	userDriver := c.Locals("user_driver").(object.Driver)

	inputUser := user.User{}
	if err := c.BodyParser(&inputUser); err != nil {
		slog.Warn(err)
		return fiber.ErrBadRequest
	}
	err := userDriver.SetType(inputUser)
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

	(*newUser)["password"] = ""
	return c.JSON(*newUser)
}

func getUser(c *fiber.Ctx) error {
	return nil
}

func setUser(c *fiber.Ctx) error {
	return nil
}

func deleteUser(c *fiber.Ctx) error {
	return nil
}
