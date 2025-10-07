package db

import (
	com "mainframe-lib/common/db"
	usr "mainframe-lib/user/model"
	"mainframe/user/config"

	"go.mongodb.org/mongo-driver/bson"
)

func AddAccount(cfg config.DBConfig, abi string, userId string, account usr.Account) error {
	// Setup timeout
	ctx, cancel := com.GetContextFromConfig(cfg.DBConfig)
	defer cancel()

	// Retrieve the collection
	coll, err := com.GetCollection(ctx, cfg.DBConfig, abi, cfg.Collections.Users)
	if err != nil {
		return err
	}

	// Setup filter
	filter := bson.M{"_id": userId}

	// Setup update command
	update := bson.M{"$addToSet": bson.M{"accounts": account}}

	// Update a document
	_, err = coll.UpdateOne(ctx, filter, update)

	return err
}

func RemoveAccount(cfg config.DBConfig, abi string, userId string, accountId usr.AccountId) error {
	// Setup timeout
	ctx, cancel := com.GetContextFromConfig(cfg.DBConfig)
	defer cancel()

	// Retrieve the collection
	coll, err := com.GetCollection(ctx, cfg.DBConfig, abi, cfg.Collections.Users)
	if err != nil {
		return err
	}

	// Setup filter
	filter := bson.M{"_id": userId}

	// Setup update command
	update := bson.M{"$pull": bson.M{"accounts": bson.M{"id": accountId}}}

	// Update a document
	_, err = coll.UpdateOne(ctx, filter, update)

	return err
}
