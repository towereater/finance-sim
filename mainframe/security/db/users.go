package db

import (
	com "mainframe-lib/common/db"
	sec "mainframe-lib/security/model"
	"mainframe/security/config"

	"go.mongodb.org/mongo-driver/bson"
)

func SelectUserByApiKey(cfg config.Config, abi string, apiKey string) (sec.User, error) {
	// Setup timeout
	ctx, cancel := com.GetContextFromConfig(cfg.DB)
	defer cancel()

	// Retrieve the collection
	coll, err := com.GetCollection(ctx, cfg.DB, abi, cfg.Prefix, cfg.Collections.Users)
	if err != nil {
		return sec.User{}, err
	}

	// Search for a document
	var user sec.User
	err = coll.FindOne(ctx, bson.M{"apiKey": apiKey}).Decode(&user)

	return user, err
}

func InsertUser(cfg config.Config, abi string, user sec.User) error {
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
