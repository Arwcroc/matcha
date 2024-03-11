package photo

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"matcha/backend/pkg/database"
	"matcha/backend/pkg/decorators/params"
	"matcha/backend/pkg/decorators/permissions"
	"matcha/backend/pkg/middleware/photoService"
	"matcha/backend/pkg/middleware/userPhotoService"
	"matcha/backend/pkg/middleware/userService"
	"matcha/backend/pkg/object"
	"matcha/backend/pkg/object/photo"
	"matcha/backend/pkg/slog"
	"matcha/backend/pkg/store"
)

func Register(app *fiber.App) {
	group := app.Group("/photo")
	group.Use(userService.UserService)
	group.Use(photoService.PhotoService)
	group.Use(userPhotoService.UserPhotoService)

	group.Get("/:username",
		permissions.LoggedIn{}.Decorate(
			params.User{}.Decorate(getPhotosOfUser).GetHandler(),
		).GetHandler(),
	)
	group.Post("/:username",
		permissions.SelfOrAdmin{}.Decorate(
			params.User{}.Decorate(createPhotoOfUser).GetHandler(),
		).GetHandler(),
	)
	group.Put("/:index", permissions.Self{}.Decorate(setPhoto).GetHandler())
	group.Delete("/:index", permissions.Self{}.Decorate(deletePhoto).GetHandler())
}

func getPhotosOfUser(c *fiber.Ctx) error {
	paramUser := c.Locals("param_user").(object.Driver)
	photoDriver := c.Locals("photo_driver").(object.Driver)
	userPhotoDriver := c.Locals("user_photo_driver").(object.Driver)

	rels, err := userPhotoDriver.GetAll(map[string]interface{}{
		"_from": paramUser.GetField("_id"),
	})
	if err != nil {
		slog.Error(err)
		return fiber.ErrInternalServerError
	}

	userPhotos := []map[string]interface{}{{}}
	for _, rel := range rels {
		userPhoto, err := photoDriver.Get(map[string]interface{}{
			"_id": rel["_to"],
		})
		if err != nil {
			if errors.Is(err, database.NotFoundError) {
				return fiber.ErrNotFound
			}
			slog.Error(err)
			return fiber.ErrInternalServerError
		}
		userPhotos = append(userPhotos, *userPhoto)
	}

	return c.JSON(userPhotos)
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

	rels, err := userPhotoDriver.GetAll(map[string]interface{}{
		"_from": paramUser.GetField("_id"),
		"_to":   photoDriver.GetField("_id"),
	})
	if err != nil {
		slog.Error(err)
		return fiber.ErrInternalServerError
	}
	if len(rels) >= 5 {
		return fiber.ErrTooManyRequests
	}

	highestIndex := -1
	for _, rel := range rels {
		index := rel["index"].(int)
		if index > highestIndex {
			highestIndex = index
		}
	}

	userPhotoDriver.SetField("_from", paramUser.GetField("_id"))
	userPhotoDriver.SetField("_to", photoDriver.GetField("_id"))
	userPhotoDriver.SetField("index", highestIndex+1)

	_, err = userPhotoDriver.Create()
	if err != nil {
		slog.Error(err)
		return fiber.ErrInternalServerError
	}

	return c.JSON(*photoDriver.GetInternal())
}

func setPhoto(c *fiber.Ctx) error {
	c.Accepts("json")

	photoDriver := c.Locals("photo_driver").(object.Driver)
	userPhotoDriver := c.Locals("user_photo_driver").(object.Driver)
	userDriver := c.Locals("user_driver").(object.Driver)
	session := c.Locals("session").(*store.Session)
	index, err := c.ParamsInt("index")
	if err != nil {
		slog.Warn(err)
		return fiber.ErrBadRequest
	}

	inputPhoto := photo.Photo{}
	err = c.BodyParser(&inputPhoto)
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

	_, err = userPhotoDriver.Get(map[string]interface{}{
		"_from": userDriver.GetField("_id"),
		"index": index,
	})
	if err != nil {
		if errors.Is(err, database.NotFoundError) {
			return fiber.ErrNotFound
		}
		slog.Error(err)
		return fiber.ErrInternalServerError
	}

	_, err = photoDriver.Get(map[string]interface{}{
		"_id": userPhotoDriver.GetField("_to"),
	})
	if err != nil {
		if errors.Is(err, database.NotFoundError) {
			return fiber.ErrNotFound
		}
		slog.Error(err)
		return fiber.ErrInternalServerError
	}

	photoDriver.SetField("b64", inputPhoto.B64)
	newPhoto, err := photoDriver.Set()
	if err != nil {
		slog.Error(err)
		return fiber.ErrInternalServerError
	}

	return c.JSON(*newPhoto)
}

func deletePhoto(c *fiber.Ctx) error {
	photoDriver := c.Locals("photo_driver").(object.Driver)
	userPhotoDriver := c.Locals("user_photo_driver").(object.Driver)
	userDriver := c.Locals("user_driver").(object.Driver)
	session := c.Locals("session").(*store.Session)
	index, err := c.ParamsInt("index")
	if err != nil {
		slog.Warn(err)
		return fiber.ErrBadRequest
	}

	inputPhoto := photo.Photo{}
	err = c.BodyParser(&inputPhoto)
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

	_, err = userPhotoDriver.Get(map[string]interface{}{
		"_from": userDriver.GetField("_id"),
		"index": index,
	})
	if err != nil {
		if errors.Is(err, database.NotFoundError) {
			return fiber.ErrNotFound
		}
		slog.Error(err)
		return fiber.ErrInternalServerError
	}

	_, err = photoDriver.Get(map[string]interface{}{
		"_id": userPhotoDriver.GetField("_to"),
	})
	if err != nil {
		if errors.Is(err, database.NotFoundError) {
			return fiber.ErrNotFound
		}
		slog.Error(err)
		return fiber.ErrInternalServerError
	}

	err = userPhotoDriver.Delete()
	if err != nil {
		slog.Error(err)
		return fiber.ErrInternalServerError
	}
	err = photoDriver.Delete()
	if err != nil {
		slog.Error(err)
		return fiber.ErrInternalServerError
	}

	return c.SendStatus(fiber.StatusOK)
}
