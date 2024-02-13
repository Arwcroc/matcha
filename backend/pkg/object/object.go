package object

type Object interface {
	Name() string
	AsMap() (map[string]interface{}, error)
}

type Driver interface {
	Create() (*map[string]interface{}, error)
	Set() (*map[string]interface{}, error)
	Get(key string, value interface{}) (*map[string]interface{}, error)
	Delete() error
	GetInternal() *map[string]interface{}
	SetType(Object) error
}
