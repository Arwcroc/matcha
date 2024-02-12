package objects

type Object interface {
	Name() string
}

type ObjectDriver interface {
	Create() (*Object, error)
	Set() (*Object, error)
	Get(key string) (*Object, error)
	Delete() error
}
