package user

import (
	"context"
	"github.com/arangodb/go-driver"
	"github.com/gofiber/fiber/v2"
	"matcha/backend/pkg/database/arangodb"
	"matcha/backend/pkg/objects/user"
	"matcha/backend/pkg/slog"
)

func getCollection(c *fiber.Ctx) error {
	arango := c.Locals("database").(*arangodb.DatabaseDriver)
	collection, err := arango.Collection(user.User{}.Name())
	if err != nil {
		slog.Error(err)
		return fiber.ErrInternalServerError
	}

	if c.Locals("database_context", arango.Ctx) == nil {
		slog.Error("could not set database_context")
		return fiber.ErrInternalServerError
	}

	if c.Locals("user_collection", collection) == nil {
		slog.Error("could not set user_collection")
		return fiber.ErrInternalServerError
	}

	return c.Next()
}

func Register(app *fiber.App) {
	group := app.Group("/user")
	group.Use(getCollection)
	group.Post("/", createUser)
	group.Get("/:id", getUser)
	group.Put("/:id", setUser)
	group.Delete("/:id", deleteUser)
}

func createUser(c *fiber.Ctx) error {
	collection := c.Locals("user_collection").(driver.Collection)
	databaseCtx := c.Locals("database_context").(context.Context)

	inputUser := user.User{}
	if err := c.BodyParser(&inputUser); err != nil {
		slog.Warn(err)
		return fiber.ErrBadRequest
	}

	userDriver := arangodb.NewObjectDriver(inputUser, collection, databaseCtx)
	newUser, err := userDriver.Create()
	if err != nil {
		slog.Error(err)
		return fiber.ErrInternalServerError
	}

	return c.JSON(newUser)
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
