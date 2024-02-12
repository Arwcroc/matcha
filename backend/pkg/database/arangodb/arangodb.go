package arangodb

import (
	"context"
	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
	"matcha/backend/pkg/database"
	"matcha/backend/pkg/slog"
)

type DatabaseDriver struct {
	database.DatabaseDriver
	Client       driver.Client
	Ctx          context.Context
	url          string
	database     *driver.Database
	databaseName string
	conn         driver.Connection
	auth         driver.Authentication
}

func New(url string, database string, auth driver.Authentication) DatabaseDriver {
	return DatabaseDriver{
		url:          url,
		auth:         auth,
		databaseName: database,
		Ctx:          context.Background(),
	}
}

func (d *DatabaseDriver) Connect() error {
	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{d.url},
	})
	if err != nil {
		return err
	}
	d.conn = conn

	client, err := driver.NewClient(driver.ClientConfig{
		Connection:     d.conn,
		Authentication: d.auth,
	})
	d.Client = client
	return err
}

func (d *DatabaseDriver) Disconnect() error {
	return nil
}

func (d *DatabaseDriver) Database() (*driver.Database, error) {
	if d.database != nil {
		return d.database, nil
	}

	exists, err := d.Client.DatabaseExists(d.Ctx, "matcha")
	if err != nil {
		return nil, err
	}
	var db driver.Database
	if !exists {
		slog.Info("Creating database matcha")
		db, err = d.Client.CreateDatabase(d.Ctx, "matcha", &driver.CreateDatabaseOptions{
			Users: []driver.CreateDatabaseUserOptions{{UserName: "matcha", Password: "matcha"}},
		})
		if err != nil {
			return nil, err
		}
	} else {
		slog.Info("Found database matcha")
		db, err = d.Client.Database(d.Ctx, "matcha")
		if err != nil {
			return nil, err
		}
	}

	collections := []string{"users"}
	var retErr error = nil
	for _, name := range collections {
		exists, err := db.CollectionExists(d.Ctx, name)
		if err != nil {
			retErr = err
			continue
		}
		if exists {
			slog.Info("Found collection", name)
			continue
		}

		slog.Info("Creating collection", name)
		_, err = db.CreateCollection(d.Ctx, name, nil)
		if err != nil {
			retErr = err
			continue
		}
	}

	d.database = &db
	return d.database, retErr
}

func (d *DatabaseDriver) Collection(name string) (driver.Collection, error) {
	db, err := d.Database()
	if err != nil {
		return nil, err
	}
	return (*db).Collection(d.Ctx, name)
}
