package object

type Object interface {
	Name() string
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
	SetInternal(Object) error
}
