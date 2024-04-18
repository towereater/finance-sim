package db

import (
	"context"
	"time"

	cfg "mainframe/user/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getCollection(db, coll string) (*mongo.Collection, error) {
	// Timeout setup
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(cfg.AppConfig.DB.Timeout)*time.Second)
	defer cancel()

	// Connection to the db
	client, err := mongo.Connect(ctx,
		options.Client().ApplyURI("mongodb://"+cfg.AppConfig.DB.Host+":"+cfg.AppConfig.DB.Port))
	if err != nil {
		return nil, err
	}

	// Check db status
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	// Retrieve the collection
	return client.Database(db).Collection(coll), nil
}
