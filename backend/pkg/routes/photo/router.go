package photo

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"matcha/backend/pkg/database"
	"matcha/backend/pkg/decorators"
	"matcha/backend/pkg/decorators/params"
	"matcha/backend/pkg/decorators/permissions"
	"matcha/backend/pkg/decorators/services"
	"matcha/backend/pkg/middleware/photoService"
	"matcha/backend/pkg/middleware/userPhotoService"
	"matcha/backend/pkg/middleware/userService"
	"matcha/backend/pkg/object/photo"
	"matcha/backend/pkg/object/user"
	"matcha/backend/pkg/object/user_photo"
	"matcha/backend/pkg/slog"
)

func Register(app *fiber.App) {
	group := app.Group("/photo")
	group.Use(userService.UserService)
	group.Use(photoService.PhotoService)
	group.Use(userPhotoService.UserPhotoService)

	group.Get("/:username", decorators.Decorate(
		getPhotosOfUser,
		permissions.LoggedIn{},
		params.User{},
	))
	group.Post("/:username", decorators.Decorate(
		createPhotoOfUser,
		permissions.SelfOrAdmin{},
		params.User{},
	))
	group.Put("/:index", decorators.Decorate(
		setPhoto,
		permissions.Self{},
		services.SelfUser{},
	))
	group.Delete("/:index", decorators.Decorate(
		deletePhoto,
		permissions.Self{},
		services.SelfUser{},
	))
}

func getPhotosOfUser(c *fiber.Ctx) error {
	paramUser := c.Locals(params.UserLocal).(user.User)
	photoObject := c.Locals(photoService.Local).(photo.Photo)
	userPhotoObject := c.Locals(userPhotoService.Local).(user_photo.UserPhoto)

	rels, err := userPhotoObject.GetAll(map[string]interface{}{
		"_from": paramUser.Id,
	})
	if err != nil {
		slog.Error(err)
		return fiber.ErrInternalServerError
	}

	var userPhotos []photo.Photo
	for _, rel := range rels {
		userPhoto, err := photoObject.Get(map[string]interface{}{
			"_id": rel.(user_photo.UserPhoto).To,
		})
		if err != nil {
			if errors.Is(err, database.NotFoundError) {
				return fiber.ErrNotFound
			}
			slog.Error(err)
			return fiber.ErrInternalServerError
		}
		userPhotos = append(userPhotos, userPhoto.(photo.Photo))
	}

	return c.JSON(userPhotos)
}

// createPhotoOfUser creates a photo with the provided json and creates a relationship with the parameter user
func createPhotoOfUser(c *fiber.Ctx) error {
	c.Accepts("json")

	paramUser := c.Locals(params.UserLocal).(user.User)
	photoObject := c.Locals(photoService.Local).(photo.Photo)
	userPhotoObject := c.Locals(userPhotoService.Local).(user_photo.UserPhoto)

	inputPhoto := photoObject
	err := c.BodyParser(&photoObject)
	if err != nil {
		slog.Warn(err)
		return fiber.ErrBadRequest
	}

	o, err := inputPhoto.Create()
	if err != nil {
		slog.Error(err)
		return fiber.ErrInternalServerError
	}
	dbPhoto := o.(photo.Photo)

	rels, err := userPhotoObject.GetAll(map[string]interface{}{
		"_from": paramUser.Id,
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
		index := rel.(user_photo.UserPhoto).Index
		if index > highestIndex {
			highestIndex = index
		}
	}

	userPhoto := userPhotoObject
	userPhoto.From = paramUser.Id
	userPhoto.To = dbPhoto.Id
	userPhoto.Index = highestIndex + 1

	_, err = userPhoto.Create()
	if err != nil {
		slog.Error(err)
		return fiber.ErrInternalServerError
	}

	return c.JSON(dbPhoto)
}

func setPhoto(c *fiber.Ctx) error {
	c.Accepts("json")

	photoObject := c.Locals(photoService.Local).(photo.Photo)
	userPhotoObject := c.Locals(userPhotoService.Local).(user_photo.UserPhoto)
	selfUser := c.Locals(services.SelfUserLocal).(user.User)

	index, err := c.ParamsInt("index")
	if err != nil {
		slog.Warn(err)
		return fiber.ErrBadRequest
	}

	inputPhoto := photoObject
	err = c.BodyParser(&inputPhoto)
	if err != nil {
		slog.Warn(err)
		return fiber.ErrBadRequest
	}

	o, err := userPhotoObject.Get(map[string]interface{}{
		"_from": selfUser.Id,
		"index": index,
	})
	if err != nil {
		if errors.Is(err, database.NotFoundError) {
			return fiber.ErrNotFound
		}
		slog.Error(err)
		return fiber.ErrInternalServerError
	}
	dbRelationship := o.(user_photo.UserPhoto)

	o, err = photoObject.Get(map[string]interface{}{
		"_id": dbRelationship.To,
	})
	if err != nil {
		if errors.Is(err, database.NotFoundError) {
			return fiber.ErrNotFound
		}
		slog.Error(err)
		return fiber.ErrInternalServerError
	}
	dbPhoto := o.(photo.Photo)

	dbPhoto.B64 = inputPhoto.B64
	o, err = dbPhoto.Set()
	if err != nil {
		slog.Error(err)
		return fiber.ErrInternalServerError
	}
	dbPhoto = o.(photo.Photo)

	return c.JSON(dbPhoto)
}

func deletePhoto(c *fiber.Ctx) error {
	photoObject := c.Locals(photoService.Local).(photo.Photo)
	userPhotoObject := c.Locals(userPhotoService.Local).(user_photo.UserPhoto)
	selfUser := c.Locals(services.SelfUserLocal).(user.User)

	index, err := c.ParamsInt("index")
	if err != nil {
		slog.Warn(err)
		return fiber.ErrBadRequest
	}

	inputPhoto := photoObject
	err = c.BodyParser(&inputPhoto)
	if err != nil {
		slog.Warn(err)
		return fiber.ErrBadRequest
	}

	o, err := userPhotoObject.Get(map[string]interface{}{
		"_from": selfUser.Id,
		"index": index,
	})
	if err != nil {
		if errors.Is(err, database.NotFoundError) {
			return fiber.ErrNotFound
		}
		slog.Error(err)
		return fiber.ErrInternalServerError
	}
	dbUserPhoto := o.(user_photo.UserPhoto)

	o, err = photoObject.Get(map[string]interface{}{
		"_id": dbUserPhoto.To,
	})
	if err != nil {
		if errors.Is(err, database.NotFoundError) {
			return fiber.ErrNotFound
		}
		slog.Error(err)
		return fiber.ErrInternalServerError
	}
	dbPhoto := o.(photo.Photo)

	err = dbUserPhoto.Delete()
	if err != nil {
		slog.Error(err)
		return fiber.ErrInternalServerError
	}
	err = dbPhoto.Delete()
	if err != nil {
		slog.Error(err)
		return fiber.ErrInternalServerError
	}

	return c.SendStatus(fiber.StatusOK)
}
