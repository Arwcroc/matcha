package user_photo

import (
	"matcha/backend/pkg/database"
	"matcha/backend/pkg/database/arangodb"
	"matcha/backend/pkg/object"
)

type UserPhoto struct {
	arangodb.EdgeDocument
	object.Object
	Index int `json:"index"`
}

func (p UserPhoto) Name() string {
	return "user_photos"
}

func (p UserPhoto) Create() (object.IObject, error) {
	return object.Create(p)
}

func (p UserPhoto) GetAll(bindValues map[string]interface{}) ([]object.IObject, error) {
	return object.GetAll(p, bindValues)
}

func (p UserPhoto) Get(bindValues map[string]interface{}) (object.IObject, error) {
	return object.Get(p, bindValues)
}

func (p UserPhoto) Set() (object.IObject, error) {
	return object.Set(p)
}

func (p UserPhoto) Delete() error {
	return object.Delete(p)
}

func (p UserPhoto) SetDatabaseDriver(driver database.Driver) object.IObject {
	p.Object.SetDatabaseDriver(driver)
	return p
}

func (p UserPhoto) SetObjectDriver(driver object.Driver) object.IObject {
	p.Object.SetObjectDriver(driver)
	return p
}

func (p UserPhoto) GetDatabaseDriver() database.Driver {
	return p.Object.GetDatabaseDriver()
}

func (p UserPhoto) GetObjectDriver() object.Driver {
	return p.Object.GetObjectDriver()
}

func (p UserPhoto) AsMap() (map[string]interface{}, error) {
	return object.AsMap(p)
}
