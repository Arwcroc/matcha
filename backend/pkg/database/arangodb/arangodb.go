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

type vertexConstraints struct {
	name              string
	vertexConstraints driver.VertexConstraints
}

type graphOptions struct {
	name              string
	vertexConstraints []vertexConstraints
	hashIndexOptions  []ensureHashIndexOptions
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

func (d *Driver) NewObjectDriver(objectType object.IObject) (object.Driver, error) {
	collection, err := d.Collection(objectType.Name())
	if err != nil {
		return nil, err
	}
	objectAsMap, err := objectType.AsMap()
	if err != nil {
		return nil, err
	}

	return &ObjectDriver{
		Ctx:           d.Ctx,
		Collection:    collection,
		wrappedType:   objectType,
		wrappedObject: objectAsMap,
	}, err
}

func (d *Driver) createCollections() {
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
		{
			"photos",
			[]ensureHashIndexOptions{},
		},
	}
	db := *d.database
	for _, collectionOption := range collections {
		exists, err := db.CollectionExists(d.Ctx, collectionOption.name)
		if err != nil {
			slog.Error(err)
			continue
		}
		if exists {
			slog.Info("Found Collection", collectionOption.name)
			continue
		}

		slog.Info("Creating Collection", collectionOption.name)
		collection, err := db.CreateCollection(d.Ctx, collectionOption.name, nil)
		if err != nil {
			slog.Error(err)
			continue
		}
		for _, options := range collectionOption.hashIndexOptions {
			_, _, err = collection.EnsureHashIndex(d.Ctx, options.fields, &options.options)
			if err != nil {
				slog.Error(err)
				continue
			}
		}
	}
}

func (d *Driver) createGraphs() {
	graphs := []graphOptions{
		{
			"user_photos",
			[]vertexConstraints{
				{
					"user_photos",
					driver.VertexConstraints{
						From: []string{"users"},
						To:   []string{"photos"},
					},
				},
			},
			[]ensureHashIndexOptions{
				{
					[]string{"Index"},
					driver.EnsureHashIndexOptions{
						Unique: true,
						Name:   "index_uniqueness",
					},
				},
			},
		},
		{
			"user_users",
			[]vertexConstraints{
				{
					"user_users",
					driver.VertexConstraints{
						From: []string{"users"},
						To:   []string{"users"},
					},
				},
			},
			[]ensureHashIndexOptions{},
		},
	}

	db := *d.database
	for _, graphOption := range graphs {
		exists, err := db.GraphExists(d.Ctx, graphOption.name)
		if err != nil {
			slog.Error(err)
			continue
		}
		if exists {
			slog.Info("Found graph", graphOption.name)
			continue
		}

		slog.Info("Creating graph", graphOption.name)
		graph, err := db.CreateGraphV2(d.Ctx, graphOption.name, nil)
		if err != nil {
			slog.Error(err)
			continue
		}

		for _, vertexConstraint := range graphOption.vertexConstraints {
			_, err := graph.CreateEdgeCollection(d.Ctx, vertexConstraint.name, vertexConstraint.vertexConstraints)
			if err != nil {
				slog.Error(err)
				continue
			}
		}

	}
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

	d.database = &db
	d.createCollections()
	d.createGraphs()

	return d.database, nil
}

func (d *Driver) Collection(name string) (driver.Collection, error) {
	db, err := d.Database()
	if err != nil {
		return nil, err
	}
	return (*db).Collection(d.Ctx, name)
}
