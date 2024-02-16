package relationship

import (
	"github.com/gofiber/fiber/v2"
	"matcha/backend/pkg/database"
	"matcha/backend/pkg/database/arangodb"
	"matcha/backend/pkg/decorators/permissions"
	"matcha/backend/pkg/object"
	"matcha/backend/pkg/object/user"
	"matcha/backend/pkg/object/user_user"
	"matcha/backend/pkg/routes"
	"matcha/backend/pkg/slog"
	"matcha/backend/pkg/store"
)

// TODO find a way to put this in the database manager middleware (through decorators ?)
func getObjectDriver(c *fiber.Ctx) error {
	dbDriver := c.Locals("database").(database.Driver)

	userDriver, err := dbDriver.NewObjectDriver(user.User{})
	if err != nil {
		return fiber.ErrInternalServerError
	}
	if c.Locals("user_driver", userDriver) == nil {
		slog.Error("could not set user_driver")
		return fiber.ErrInternalServerError
	}

	userUserDriver, err := dbDriver.NewObjectDriver(user_user.UserUser{})
	if err != nil {
		slog.Error(err)
		return fiber.ErrInternalServerError
	}
	if c.Locals("user_user_driver", userUserDriver) == nil {
		slog.Error("could not set user_user_driver")
		return fiber.ErrInternalServerError
	}

	return c.Next()
}

func Register(app *fiber.App) {
	group := app.Group("/relationship")
	group.Use(getObjectDriver)
	group.Put("/:username", permissions.LoggedIn(routes.GetUserFromParam(setRelationship)))
}

// TODO make this generic
// setRelationship Updates or creates a relationship element between two users with the provided json
func setRelationship(c *fiber.Ctx) error {
	c.Accepts("json")
	userUserDriver := c.Locals("user_user_driver").(arangodb.ObjectDriver)
	userDriver := c.Locals("user_driver").(object.Driver)
	paramUser := c.Locals("param_user").(object.Driver)
	session := c.Locals("session").(*store.Session)

	inputRelationship := user_user.UserUser{}
	err := c.BodyParser(&inputRelationship)
	if err != nil {
		slog.Warn(err)
		return fiber.ErrBadRequest
	}

	_, err = userDriver.Get(map[string]interface{}{
		"username": session.Get("username").(string),
	})
	if err != nil {
		slog.Error(err)
		return fiber.ErrInternalServerError
	}

	_, err = userUserDriver.Get(map[string]interface{}{
		"_from": userDriver.GetField("_id"),
		"_to":   paramUser.GetField("_id"),
	})
	if err != nil {
		slog.Error(err)
		return fiber.ErrBadRequest
	}

	err = userUserDriver.SetInternal(inputRelationship)
	if err != nil {
		slog.Warn(err)
		return fiber.ErrBadRequest
	}

	userUserDriver.SetField("_from", userDriver.GetField("_id"))
	userUserDriver.SetField("_to", paramUser.GetField("_id"))

	relationship, err := userUserDriver.Create()
	if err != nil {
		slog.Error(err)
		return fiber.ErrInternalServerError
	}

	return c.JSON(*relationship)
}
