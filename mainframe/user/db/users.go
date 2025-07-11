package db

import (
	"mainframe/user/config"
	"mainframe/user/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SelectUser(cfg config.Config, abi string, userId primitive.ObjectID) (model.User, error) {
	// Setup timeout
	ctx, cancel := getContextFromConfig(cfg.DB)
	defer cancel()

	// Retrieve the collection
	coll, err := getCollection(ctx, cfg.DB, abi, cfg.Prefix, cfg.Collections.Users)
	if err != nil {
		return model.User{}, err
	}

	// Search for a document
	var user model.User
	err = coll.FindOne(ctx, bson.M{"_id": userId}).Decode(&user)

	return user, err
}

func SelectUsers(cfg config.Config, abi string, userFilter model.User, from primitive.ObjectID, limit int) ([]model.User, error) {
	// Setup timeout
	ctx, cancel := getContextFromConfig(cfg.DB)
	defer cancel()

	// Retrieve the collection
	coll, err := getCollection(ctx, cfg.DB, abi, cfg.Prefix, cfg.Collections.Users)
	if err != nil {
		return []model.User{}, err
	}

	// Setup find options
	var opts options.FindOptions
	opts.SetLimit(int64(limit))

	// Setup filter
	filter := bson.M{}
	if userFilter.Username != "" {
		filter["username"] = userFilter.Username
	}
	filter["_id"] = bson.M{"$gt": from}

	// Define the cursor
	cursor, err := coll.Find(ctx, filter, &opts)
	if err != nil {
		return []model.User{}, err
	}

	// Search for the documents
	var users []model.User
	err = cursor.All(ctx, &users)
	if err != nil {
		return []model.User{}, err
	}

	if users == nil {
		return []model.User{}, nil
	}
	return users, err
}

func InsertUser(cfg config.Config, abi string, user model.User) error {
	// Setup timeout
	ctx, cancel := getContextFromConfig(cfg.DB)
	defer cancel()

	// Retrieve the collection
	coll, err := getCollection(ctx, cfg.DB, abi, cfg.Prefix, cfg.Collections.Users)
	if err != nil {
		return err
	}

	// Insert a document
	_, err = coll.InsertOne(ctx, user)

	return err
}

func UpdateUser(cfg config.Config, abi string, userId primitive.ObjectID, user model.User) error {
	// Setup timeout
	ctx, cancel := getContextFromConfig(cfg.DB)
	defer cancel()

	// Retrieve the collection
	coll, err := getCollection(ctx, cfg.DB, abi, cfg.Prefix, cfg.Collections.Users)
	if err != nil {
		return err
	}

	// Setup filter
	filter := bson.M{"_id": userId}

	// Setup update command
	update := bson.M{"$set": bson.M{
		"username": user.Username,
		"password": user.Password,
		"name":     user.Name,
		"surname":  user.Surname,
		"birth":    user.Birth,
	}}

	// Update a document
	_, err = coll.UpdateOne(ctx, filter, update)

	return err
}

func DeleteUser(cfg config.Config, abi string, userId primitive.ObjectID) error {
	// Setup timeout
	ctx, cancel := getContextFromConfig(cfg.DB)
	defer cancel()

	// Retrieve the collection
	coll, err := getCollection(ctx, cfg.DB, abi, cfg.Prefix, cfg.Collections.Users)
	if err != nil {
		return err
	}

	// Delete a document
	_, err = coll.DeleteOne(ctx, bson.M{"_id": userId})

	return err
}
