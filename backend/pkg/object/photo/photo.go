package photo

import (
	"matcha/backend/pkg/database"
	"matcha/backend/pkg/database/arangodb"
	"matcha/backend/pkg/object"
)

type Photo struct {
	arangodb.Document
	object.Object
	B64 string `json:"b64"`
}

func (p Photo) Name() string {
	return "photos"
}

func (p Photo) Create() (object.IObject, error) {
	return object.Create(p)
}

func (p Photo) Get(bindValues map[string]interface{}) (object.IObject, error) {
	return object.Get(p, bindValues)
}

func (p Photo) GetAll(bindValues map[string]interface{}) ([]object.IObject, error) {
	return object.GetAll(p, bindValues)
}

func (p Photo) Set() (object.IObject, error) {
	return object.Set(p)
}

func (p Photo) Delete() error {
	return object.Delete(p)
}

func (p Photo) SetDatabaseDriver(driver database.Driver) object.IObject {
	p.Object.SetDatabaseDriver(driver)
	return p
}

func (p Photo) SetObjectDriver(driver object.Driver) object.IObject {
	p.Object.SetObjectDriver(driver)
	return p
}

func (p Photo) GetDatabaseDriver() database.Driver {
	return p.Object.GetDatabaseDriver()
}

func (p Photo) GetObjectDriver() object.Driver {
	return p.Object.GetObjectDriver()
}

func (p Photo) AsMap() (map[string]interface{}, error) {
	return object.AsMap(p)
}
