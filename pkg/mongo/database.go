package mongo

import (
	"context"

	mongodb "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DatabaseConfig struct {
	ConnectionString string
	Database         string
}
type Database struct {
	Config DatabaseConfig
	client mongodb.Client
}

func (db *Database) GetCollection(dbName string, collectionName string) mongodb.Collection {
	return *db.client.Database(dbName).Collection(collectionName)
}

func Connect(context context.Context, config DatabaseConfig) Database {

	client, err := mongodb.Connect(context, options.Client().ApplyURI(config.ConnectionString))
	if err != nil {
		panic(err)
	}

	return Database{
		Config: config,
		client: *client,
	}
}
