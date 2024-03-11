package relationship

import (
	"github.com/gofiber/fiber/v2"
	"matcha/backend/pkg/database/arangodb"
	"matcha/backend/pkg/decorators/params"
	"matcha/backend/pkg/decorators/permissions"
	"matcha/backend/pkg/middleware/userService"
	"matcha/backend/pkg/middleware/userUserService"
	"matcha/backend/pkg/object"
	"matcha/backend/pkg/object/user_user"
	"matcha/backend/pkg/slog"
	"matcha/backend/pkg/store"
)

func Register(app *fiber.App) {
	group := app.Group("/relationship")
	group.Use(userService.UserService)
	group.Use(userUserService.UserUserService)

	group.Put("/:username",
		permissions.LoggedIn{}.Decorate(
			params.User{}.Decorate(setRelationship).GetHandler(),
		).GetHandler(),
	)
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
