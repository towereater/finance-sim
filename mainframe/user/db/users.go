package db

import (
	com "mainframe-lib/common/db"
	usr "mainframe-lib/user/model"
	"mainframe/user/config"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SelectUser(cfg config.Config, abi string, userId string) (usr.User, error) {
	// Setup timeout
	ctx, cancel := com.GetContextFromConfig(cfg.DB)
	defer cancel()

	// Retrieve the collection
	coll, err := com.GetCollection(ctx, cfg.DB, abi, cfg.Prefix, cfg.Collections.Users)
	if err != nil {
		return usr.User{}, err
	}

	// Search for a document
	var user usr.User
	err = coll.FindOne(ctx, bson.M{"_id": userId}).Decode(&user)

	return user, err
}

func SelectUsers(cfg config.Config, abi string, userFilter usr.User, from string, limit int) ([]usr.User, error) {
	// Setup timeout
	ctx, cancel := com.GetContextFromConfig(cfg.DB)
	defer cancel()

	// Retrieve the collection
	coll, err := com.GetCollection(ctx, cfg.DB, abi, cfg.Prefix, cfg.Collections.Users)
	if err != nil {
		return []usr.User{}, err
	}

	// Setup find options
	var opts options.FindOptions
	opts.SetLimit(int64(limit))

	// Setup filter
	filter := bson.M{}
	if userFilter.Username != "" {
		filter["username"] = userFilter.Username
	}
	if userFilter.Password != "" {
		filter["password"] = userFilter.Password
	}
	filter["_id"] = bson.M{"$gt": from}

	// Define the cursor
	cursor, err := coll.Find(ctx, filter, &opts)
	if err != nil {
		return []usr.User{}, err
	}

	// Search for the documents
	var users []usr.User
	err = cursor.All(ctx, &users)
	if err != nil {
		return []usr.User{}, err
	}

	if users == nil {
		return []usr.User{}, nil
	}
	return users, err
}

func InsertUser(cfg config.Config, abi string, user usr.User) error {
	// Setup timeout
	ctx, cancel := com.GetContextFromConfig(cfg.DB)
	defer cancel()

	// Retrieve the collection
	coll, err := com.GetCollection(ctx, cfg.DB, abi, cfg.Prefix, cfg.Collections.Users)
	if err != nil {
		return err
	}

	// Insert a document
	_, err = coll.InsertOne(ctx, user)

	return err
}

func UpdateUser(cfg config.Config, abi string, userId string, user usr.User) error {
	// Setup timeout
	ctx, cancel := com.GetContextFromConfig(cfg.DB)
	defer cancel()

	// Retrieve the collection
	coll, err := com.GetCollection(ctx, cfg.DB, abi, cfg.Prefix, cfg.Collections.Users)
	if err != nil {
		return err
	}

	// Setup filter
	filter := bson.M{"_id": userId}

	// Setup update command
	update := bson.M{"$set": bson.M{
		"password": user.Password,
	}}

	// Update a document
	_, err = coll.UpdateOne(ctx, filter, update)

	return err
}

func DeleteUser(cfg config.Config, abi string, userId string) error {
	// Setup timeout
	ctx, cancel := com.GetContextFromConfig(cfg.DB)
	defer cancel()

	// Retrieve the collection
	coll, err := com.GetCollection(ctx, cfg.DB, abi, cfg.Prefix, cfg.Collections.Users)
	if err != nil {
		return err
	}

	// Delete a document
	_, err = coll.DeleteOne(ctx, bson.M{"_id": userId})

	return err
}
