package user

import (
	"matcha/backend/pkg/database"
	"matcha/backend/pkg/database/arangodb"
	"matcha/backend/pkg/object"
)

type User struct {
	arangodb.Document
	object.Object
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Password  string `json:"password,omitempty"`
}

func (u User) Name() string {
	return "users"
}

func (u User) Create() (object.IObject, error) {
	return object.Create(u)
}

func (u User) Get(bindValues map[string]interface{}) (object.IObject, error) {
	return object.Get(u, bindValues)
}

func (u User) GetAll(bindValues map[string]interface{}) ([]object.IObject, error) {
	return object.GetAll(u, bindValues)
}

func (u User) Set() (object.IObject, error) {
	return object.Set(u)
}

func (u User) Delete() error {
	return object.Delete(u)
}

func (u User) SetDatabaseDriver(driver database.Driver) object.IObject {
	u.Object.SetDatabaseDriver(driver)
	return u
}

func (u User) SetObjectDriver(driver object.Driver) object.IObject {
	u.Object.SetObjectDriver(driver)
	return u
}

func (u User) GetDatabaseDriver() database.Driver {
	return u.Object.GetDatabaseDriver()
}

func (u User) GetObjectDriver() object.Driver {
	return u.Object.GetObjectDriver()
}

func (u User) AsMap() (map[string]interface{}, error) {
	return object.AsMap(u)
}
