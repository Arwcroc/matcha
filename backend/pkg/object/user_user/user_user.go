package user_user

import (
	"matcha/backend/pkg/database"
	"matcha/backend/pkg/database/arangodb"
	"matcha/backend/pkg/object"
)

type RelationshipType string

const (
	RelationshipPass  RelationshipType = "pass"
	RelationshipSmash RelationshipType = "smash"
)

type UserUser struct {
	arangodb.EdgeDocument
	object.Object
	Relationship RelationshipType `json:"relationship"`
}

func (u UserUser) Name() string {
	return "user_users"
}

func (u UserUser) Create() (object.IObject, error) {
	return object.Create(u)
}

func (u UserUser) Get(bindValues map[string]interface{}) (object.IObject, error) {
	return object.Get(u, bindValues)
}

func (u UserUser) GetAll(bindValues map[string]interface{}) ([]object.IObject, error) {
	return object.GetAll(u, bindValues)
}

func (u UserUser) Set() (object.IObject, error) {
	return object.Set(u)
}

func (u UserUser) Delete() error {
	return object.Delete(u)
}

func (u UserUser) SetDatabaseDriver(driver database.Driver) object.IObject {
	u.Object.SetDatabaseDriver(driver)
	return u
}

func (u UserUser) SetObjectDriver(driver object.Driver) object.IObject {
	u.Object.SetObjectDriver(driver)
	return u
}

func (u UserUser) GetDatabaseDriver() database.Driver {
	return u.Object.GetDatabaseDriver()
}

func (u UserUser) GetObjectDriver() object.Driver {
	return u.Object.GetObjectDriver()
}

func (u UserUser) AsMap() (map[string]interface{}, error) {
	return object.AsMap(u)
}
