package photo

import (
	"github.com/gofiber/fiber/v2"
	"matcha/backend/pkg/database"
	"matcha/backend/pkg/decorators/permissions"
	"matcha/backend/pkg/object"
	"matcha/backend/pkg/object/photo"
	"matcha/backend/pkg/object/user"
	"matcha/backend/pkg/object/user_photo"
	"matcha/backend/pkg/routes"
	"matcha/backend/pkg/slog"
)

func getObjectDriver(c *fiber.Ctx) error {
	dbDriver := c.Locals("database").(database.Driver)

	userDriver, err := dbDriver.NewObjectDriver(user.User{})
	if err != nil {
		slog.Error(err)
		return fiber.ErrInternalServerError
	}
	if c.Locals("user_driver", userDriver) == nil {
		slog.Error("could not set user_driver")
		return fiber.ErrInternalServerError
	}

	photoDriver, err := dbDriver.NewObjectDriver(photo.Photo{})
	if err != nil {
		slog.Error(err)
		return fiber.ErrInternalServerError
	}
	if c.Locals("photo_driver", photoDriver) == nil {
		slog.Error("could not set photo_driver")
		return fiber.ErrInternalServerError
	}

	userPhotoDriver, err := dbDriver.NewObjectDriver(user_photo.UserPhoto{})
	if err != nil {
		slog.Error(err)
		return fiber.ErrInternalServerError
	}
	if c.Locals("user_photo_driver", userPhotoDriver) == nil {
		slog.Error("could not set user_photo_driver")
		return fiber.ErrInternalServerError
	}

	return c.Next()
}

func Register(app *fiber.App) {
	group := app.Group("/photo")
	group.Use(getObjectDriver)
	//group.Get("/:username", permissions.SelfOrAdmin(routes.GetUserFromParam(getPhotosOfUser)))
	group.Post("/:username", permissions.SelfOrAdmin(routes.GetUserFromParam(createPhotoOfUser)))
	//group.Put("/:id", permissions.SelfOrAdmin(setPhoto))
	//group.Delete("/:id", permissions.SelfOrAdmin(deletePhoto))
}

// TODO Find a way to make relations generic
// createPhotoOfUser creates a photo with the provided json and creates a relationship with the parameter user
func createPhotoOfUser(c *fiber.Ctx) error {
	c.Accepts("json")

	paramUser := c.Locals("param_user").(object.Driver)
	photoDriver := c.Locals("photo_driver").(object.Driver)
	userPhotoDriver := c.Locals("user_photo_driver").(object.Driver)

	inputPhoto := photo.Photo{}
	err := c.BodyParser(&inputPhoto)
	if err != nil {
		slog.Warn(err)
		return fiber.ErrBadRequest
	}

	err = photoDriver.SetInternal(inputPhoto)
	if err != nil {
		slog.Warn(err)
		return fiber.ErrBadRequest
	}

	_, err = photoDriver.Create()
	if err != nil {
		slog.Error(err)
		return fiber.ErrInternalServerError
	}

	userPhotoDriver.SetField("_from", paramUser.GetField("_id"))
	userPhotoDriver.SetField("_to", photoDriver.GetField("_id"))

	_, err = userPhotoDriver.Create()
	if err != nil {
		slog.Error(err)
		return fiber.ErrInternalServerError
	}

	return c.JSON(*photoDriver.GetInternal())
}
