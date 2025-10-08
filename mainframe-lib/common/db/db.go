package db

import (
	"context"
	"fmt"

	"mainframe-lib/common/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetCollection(ctx context.Context, db config.DB, abi string, coll string) (*mongo.Collection, error) {
	// Connect to the db
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://"+db.Host))
	if err != nil {
		return nil, err
	}

	// Retrieve the collection
	return client.Database(getAbiDBName(abi, db.Prefix)).Collection(coll), nil
}

func getAbiDBName(abi string, db string) string {
	return fmt.Sprintf("%s-%s", abi, db)
}
