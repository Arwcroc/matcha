package database

type DatabaseDriver interface {
	Connect() error
	Disconnect() error
}
