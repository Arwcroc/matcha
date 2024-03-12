package object

import (
	"encoding/json"
	"matcha/backend/pkg/database"
)

type IObject interface {
	Name() string
	Create() (IObject, error)
	Get(bindValues map[string]interface{}) (IObject, error)
	GetAll(bindValues map[string]interface{}) ([]IObject, error)
	Set() (IObject, error)
	Delete() error
	SetDatabaseDriver(driver database.Driver) IObject
	SetObjectDriver(driver Driver) IObject
	GetDatabaseDriver() database.Driver
	GetObjectDriver() Driver
	AsMap() (map[string]interface{}, error)
}

type Driver interface {
	Create() (*map[string]interface{}, error)
	Set() (*map[string]interface{}, error)
	Get(bindValues map[string]interface{}) (*map[string]interface{}, error)
	GetAll(bindValues map[string]interface{}) ([]map[string]interface{}, error)
	Delete() error
	GetField(string) interface{}
	SetField(string, interface{})
	GetInternal() *map[string]interface{}
	SetInternal(IObject) error
}

type Object struct {
	databaseDriver database.Driver
	objectDriver   Driver
}

func (o *Object) GetDatabaseDriver() database.Driver {
	return o.databaseDriver
}

func (o *Object) GetObjectDriver() Driver {
	return o.objectDriver
}

func (o *Object) SetDatabaseDriver(driver database.Driver) {
	o.databaseDriver = driver
}

func (o *Object) SetObjectDriver(driver Driver) {
	o.objectDriver = driver
}

func New[O IObject](driver database.Driver) (object O, err error) {
	objectDriver, err := driver.NewObjectDriver(object)
	object.SetDatabaseDriver(driver)
	object.SetObjectDriver(objectDriver)
	return object, err
}

func NewFromMap[O IObject](asMap map[string]interface{}, driver database.Driver) (object O, err error) {
	asBytes, err := json.Marshal(asMap)
	if err != nil {
		return object, err
	}
	object, err = New[O](driver)
	if err != nil {
		return object, err
	}
	err = json.Unmarshal(asBytes, &object)
	return object, err
}

func NewFromObjectDriver[O IObject](objectDriver Driver, driver database.Driver) (O, error) {
	return NewFromMap[O](*objectDriver.GetInternal(), driver)
}

func Create[O IObject](object O) (IObject, error) {
	err := object.GetObjectDriver().SetInternal(object)
	if err != nil {
		return nil, err
	}
	_, err = object.GetObjectDriver().Create()
	if err != nil {
		return nil, err
	}
	return NewFromObjectDriver[O](object.GetObjectDriver(), object.GetDatabaseDriver())
}

func Get[O IObject](object O, bindValues map[string]interface{}) (O, error) {
	err := object.GetObjectDriver().SetInternal(object)
	if err != nil {
		return object, err
	}
	_, err = object.GetObjectDriver().Get(bindValues)
	if err != nil {
		return object, err
	}

	return NewFromObjectDriver[O](object.GetObjectDriver(), object.GetDatabaseDriver())
}

func GetAll[O IObject](object O, bindValues map[string]interface{}) ([]IObject, error) {
	err := object.GetObjectDriver().SetInternal(object)
	if err != nil {
		return nil, err
	}
	allObjects, err := object.GetObjectDriver().GetAll(bindValues)
	if err != nil {
		return nil, err
	}

	var objects []IObject
	for _, obj := range allObjects {
		newObj, err := NewFromMap[O](obj, object.GetDatabaseDriver())
		if err != nil {
			return nil, err
		}
		objects = append(objects, newObj)
	}

	return objects, nil
}

func Set[O IObject](object O) (IObject, error) {
	err := object.GetObjectDriver().SetInternal(object)
	if err != nil {
		return nil, err
	}
	_, err = object.GetObjectDriver().Set()
	if err != nil {
		return nil, err
	}
	return NewFromObjectDriver[O](object.GetObjectDriver(), object.GetDatabaseDriver())
}

func Delete[O IObject](object O) error {
	err := object.GetObjectDriver().SetInternal(object)
	if err != nil {
		return err
	}
	return object.GetObjectDriver().Delete()
}

func AsMap[O IObject](object O) (map[string]interface{}, error) {
	asBytes, err := json.Marshal(object)
	if err != nil {
		return nil, err
	}

	asMap := map[string]interface{}{}
	err = json.Unmarshal(asBytes, &asMap)
	return asMap, err
}
