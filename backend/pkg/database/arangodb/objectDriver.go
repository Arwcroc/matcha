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

type Document struct {
	Key string `json:"_key,omitempty"`
	Id  string `json:"_id,omitempty"`
	Rev string `json:"_rev,omitempty"`
}

type EdgeDocument struct {
	Document
	From string `json:"_from,omitempty"`
	To   string `json:"_to,omitempty"`
}

type ObjectDriver struct {
	wrappedType   object.IObject
	wrappedObject map[string]interface{}
	Ctx           context.Context
	Collection    driver.Collection
}

func (o *ObjectDriver) withReturnNew() context.Context {
	return driver.WithReturnNew(o.Ctx, &o.wrappedObject)
}

func (o *ObjectDriver) Name() string {
	return o.Collection.Name()
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

func (o *ObjectDriver) SetInternal(object object.IObject) error {
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
	_, err := o.Collection.CreateDocument(o.withReturnNew(), o.wrappedObject)
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
	_, err := o.Collection.UpdateDocument(o.withReturnNew(), o.wrappedObject["_key"].(string), o.wrappedObject)
	return &o.wrappedObject, err
}

func (o *ObjectDriver) Get(bindValues map[string]interface{}) (*map[string]interface{}, error) {
	query := fmt.Sprintf("FOR doc IN %s FILTER", o.Collection.Name())
	for key := range bindValues {
		query = fmt.Sprintf("%s doc.%s == @%s", query, key, key)
	}
	query = fmt.Sprintf("%s RETURN doc", query)

	cursor, err := o.Collection.Database().Query(
		o.withReturnNew(),
		query,
		bindValues,
	)
	if err != nil {
		return nil, err
	}
	defer cursor.Close()
	if !cursor.HasMore() {
		return nil, database.NotFoundError
	}
	_, err = cursor.ReadDocument(o.Ctx, &o.wrappedObject)
	return &o.wrappedObject, err
}

func (o *ObjectDriver) GetAll(bindValues map[string]interface{}) ([]map[string]interface{}, error) {
	query := fmt.Sprintf("FOR doc IN %s FILTER", o.Collection.Name())
	for key := range bindValues {
		query = fmt.Sprintf("%s doc.%s == @%s", query, key, key)
	}
	query = fmt.Sprintf("%s RETURN doc", query)

	cursor, err := o.Collection.Database().Query(
		o.withReturnNew(),
		query,
		bindValues,
	)
	if err != nil {
		return nil, err
	}
	defer cursor.Close()

	docs := []map[string]interface{}{{}}
	for cursor.HasMore() {
		doc := map[string]interface{}{}
		_, err = cursor.ReadDocument(o.withReturnNew(), &doc)
		if err != nil {
			return nil, err
		}
		docs = append(docs, doc)
	}

	return docs, nil
}

func (o *ObjectDriver) Delete() error {
	if o.wrappedObject["_key"] == nil {
		return errors.New("object has no _key")
	}
	_, err := o.Collection.RemoveDocument(o.Ctx, o.wrappedObject["_key"].(string))
	return err
}
