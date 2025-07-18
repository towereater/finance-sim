package db

import (
	"context"
	"fmt"
	"time"

	"mainframe/account/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getCollection(ctx context.Context, cfg config.DBConfig, abi string, db string, coll string) (*mongo.Collection, error) {
	// Connect to the db
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://"+cfg.Host))
	if err != nil {
		return nil, err
	}

	// Retrieve the collection
	return client.Database(getAbiDBName(abi, db)).Collection(coll), nil
}

func getAbiDBName(abi string, db string) string {
	return fmt.Sprintf("%s-%s", abi, db)
}

func getContextFromConfig(cfg config.DBConfig) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), time.Duration(cfg.Timeout)*time.Second)
}
