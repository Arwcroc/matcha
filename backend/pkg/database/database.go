package database

import "matcha/backend/pkg/object"

type Driver interface {
	Connect() error
	Disconnect() error
	NewObjectDriver(objectType object.Object) (object.Driver, error)
}

type Error struct {
	message string
}

func NewError(message string) *Error {
	return &Error{message: message}
}

func (e *Error) Error() string {
	return e.message
}

var (
	UniqueConstraintError = NewError("unique constraint violated")
)
