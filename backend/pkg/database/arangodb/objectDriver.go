package arangodb

import (
	"context"
	"errors"
	"fmt"
	"github.com/arangodb/go-driver"
	"matcha/backend/pkg/database"
	"matcha/backend/pkg/object"
	"matcha/backend/pkg/slog"
	"strings"
)

type ObjectDriver struct {
	wrappedType   object.Object
	wrappedObject map[string]interface{}
	ctx           context.Context
	collection    driver.Collection
}

func (o *ObjectDriver) withReturnNew() context.Context {
	return driver.WithReturnNew(o.ctx, &o.wrappedObject)
}

func (o *ObjectDriver) Name() string {
	return o.collection.Name()
}

func (o *ObjectDriver) GetField(key string) interface{} {
	return o.wrappedObject[key]
}

func (o *ObjectDriver) SetField(key string, value interface{}) {
	o.wrappedObject[key] = value
}

func (o *ObjectDriver) GetInternal() *map[string]interface{} {
	return &o.wrappedObject
}

func (o *ObjectDriver) SetInternal(object object.Object) error {
	o.wrappedType = object
	asMap, err := o.wrappedType.AsMap()
	if err != nil {
		return err
	}
	if o.wrappedObject["_key"] != nil {
		asMap["_key"] = o.wrappedObject["_key"]
	}
	o.wrappedObject = asMap
	return nil
}

func (o *ObjectDriver) Create() (*map[string]interface{}, error) {
	_, err := o.collection.CreateDocument(o.withReturnNew(), o.wrappedObject)
	if err != nil {
		if strings.Contains(err.Error(), "unique constraint violated") {
			slog.Debug(err)
			err = database.UniqueConstraintError
		}
	}
	return &o.wrappedObject, err
}

func (o *ObjectDriver) Set() (*map[string]interface{}, error) {
	if o.wrappedObject["_key"] == nil {
		return nil, errors.New("object has no _key")
	}
	_, err := o.collection.UpdateDocument(o.withReturnNew(), o.wrappedObject["_key"].(string), o.wrappedObject)
	return &o.wrappedObject, err
}

func (o *ObjectDriver) Get(key string, value interface{}) (*map[string]interface{}, error) {
	query := fmt.Sprintf("FOR doc IN %s FILTER doc.%s == @fieldValue RETURN doc", o.collection.Name(), key)

	cursor, err := o.collection.Database().Query(
		o.withReturnNew(),
		query,
		map[string]interface{}{
			"fieldValue": value,
		},
	)
	if err != nil {
		return nil, err
	}
	defer cursor.Close()
	if !cursor.HasMore() {
		return nil, database.NotFoundError
	}
	_, err = cursor.ReadDocument(o.ctx, &o.wrappedObject)
	return &o.wrappedObject, err
}

func (o *ObjectDriver) Delete() error {
	if o.wrappedObject["_key"] == nil {
		return errors.New("object has no _key")
	}
	_, err := o.collection.RemoveDocument(o.ctx, o.wrappedObject["_key"].(string))
	return err
}
