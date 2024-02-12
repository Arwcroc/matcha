package arangodb

import (
	"context"
	"github.com/arangodb/go-driver"
	"matcha/backend/pkg/objects"
)

type ObjectDriver struct {
	objects.ObjectDriver
	objects.Object
	Wrapped    objects.Object
	collection driver.Collection
	ctx        context.Context
}

func NewObjectDriver(object objects.Object, collection driver.Collection, ctx context.Context) ObjectDriver {
	return ObjectDriver{
		Wrapped:    object,
		collection: collection,
		ctx:        ctx,
	}
}

func (o *ObjectDriver) Name() string {
	return o.Wrapped.Name()
}

func (o *ObjectDriver) Create() (*objects.Object, error) {
	_, err := o.collection.CreateDocument(driver.WithReturnNew(o.ctx, &o.Wrapped), o.Wrapped)
	return &o.Wrapped, err
}

func (o *ObjectDriver) Set() (*objects.Object, error) {
	//TODO implement me
	panic("implement me")
}

func (o *ObjectDriver) Get(key string) (*objects.Object, error) {
	_, err := o.collection.ReadDocument(o.ctx, key, &o.Wrapped)
	return &o.Wrapped, err
}

func (o *ObjectDriver) Delete() error {
	//TODO implement me
	panic("implement me")
}
