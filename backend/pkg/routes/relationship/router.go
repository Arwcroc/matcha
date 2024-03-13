package relationship

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"matcha/backend/pkg/database"
	"matcha/backend/pkg/decorators"
	"matcha/backend/pkg/decorators/params"
	"matcha/backend/pkg/decorators/permissions"
	"matcha/backend/pkg/decorators/services"
	"matcha/backend/pkg/middleware/userService"
	"matcha/backend/pkg/middleware/userUserService"
	"matcha/backend/pkg/object/user"
	"matcha/backend/pkg/object/user_user"
	"matcha/backend/pkg/slog"
)

func Register(app *fiber.App) {
	group := app.Group("/relationship")
	group.Use(userService.UserService)
	group.Use(userUserService.UserUserService)

	group.Put("/:username", decorators.Decorate(
		setRelationship,
		permissions.LoggedIn{},
		params.User{},
		services.SelfUser{},
	))
}

// setRelationship Updates or creates a relationship element between two users with the provided json
func setRelationship(c *fiber.Ctx) error {
	c.Accepts("json")
	userUserObject := c.Locals(userUserService.Local).(user_user.UserUser)
	selfUser := c.Locals(services.SelfUserLocal).(user.User)
	paramUser := c.Locals(params.UserLocal).(user.User)

	inputRelationship := userUserObject
	err := c.BodyParser(&inputRelationship)
	if err != nil {
		slog.Warn(err)
		return fiber.ErrBadRequest
	}

	inputRelationship.From = selfUser.Id
	inputRelationship.To = paramUser.Id

	o, err := userUserObject.Get(map[string]interface{}{
		"_from": inputRelationship.From,
		"_to":   inputRelationship.To,
	})
	if err != nil {
		if errors.Is(err, database.NotFoundError) {
			o, err = inputRelationship.Create()
			if err != nil {
				slog.Error(err)
				return fiber.ErrInternalServerError
			}
			return c.JSON(o.(user_user.UserUser))
		}
		slog.Error(err)
		return fiber.ErrBadRequest
	}
	dbRelationship := o.(user_user.UserUser)
	inputRelationship.Key = dbRelationship.Key

	return c.JSON(inputRelationship)
}
