package arangodb

import (
	"context"
	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
	"matcha/backend/pkg/object"
	"matcha/backend/pkg/slog"
)

type Driver struct {
	Client       driver.Client
	Ctx          context.Context
	url          string
	database     *driver.Database
	databaseName string
	conn         driver.Connection
	auth         driver.Authentication
}

type ensureHashIndexOptions struct {
	fields  []string
	options driver.EnsureHashIndexOptions
}

type collectionOptions struct {
	name             string
	hashIndexOptions []ensureHashIndexOptions
}

func New(url string, database string, auth driver.Authentication) Driver {
	return Driver{
		url:          url,
		auth:         auth,
		databaseName: database,
		Ctx:          context.Background(),
	}
}

func (d *Driver) Connect() error {
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

func (d *Driver) Disconnect() error {
	return nil
}

func (d *Driver) NewObjectDriver(objectType object.Object) (object.Driver, error) {
	collection, err := d.Collection(objectType.Name())
	if err != nil {
		return nil, err
	}
	objectAsMap, err := objectType.AsMap()
	if err != nil {
		return nil, err
	}

	return &ObjectDriver{
		ctx:           d.Ctx,
		collection:    collection,
		wrappedType:   objectType,
		wrappedObject: objectAsMap,
	}, err
}

func (d *Driver) Database() (*driver.Database, error) {
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

	collections := []collectionOptions{
		{
			"users",
			[]ensureHashIndexOptions{
				{
					[]string{"username"},
					driver.EnsureHashIndexOptions{
						Unique: true,
						Name:   "username_uniqueness",
					},
				},
				{
					[]string{"email"},
					driver.EnsureHashIndexOptions{
						Unique: true,
						Name:   "email_uniqueness",
					},
				},
			},
		},
	}
	var retErr error = nil
	for _, collectionOption := range collections {
		exists, err := db.CollectionExists(d.Ctx, collectionOption.name)
		if err != nil {
			retErr = err
			continue
		}
		if exists {
			slog.Info("Found collection", collectionOption.name)
			continue
		}

		slog.Info("Creating collection", collectionOption.name)
		collection, err := db.CreateCollection(d.Ctx, collectionOption.name, nil)
		if err != nil {
			retErr = err
			continue
		}
		for _, options := range collectionOption.hashIndexOptions {
			_, _, err = collection.EnsureHashIndex(d.Ctx, options.fields, &options.options)
			if err != nil {
				retErr = err
				continue
			}
		}
	}

	d.database = &db
	return d.database, retErr
}

func (d *Driver) Collection(name string) (driver.Collection, error) {
	db, err := d.Database()
	if err != nil {
		return nil, err
	}
	return (*db).Collection(d.Ctx, name)
}
