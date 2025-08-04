package db

import (
	"mainframe/user/config"
	"mainframe/user/model"

	"go.mongodb.org/mongo-driver/bson"
)

func AddAccount(cfg config.Config, abi string, userId string, account model.Account) error {
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
	update := bson.M{"$addToSet": bson.M{"accounts": account}}

	// Update a document
	_, err = coll.UpdateOne(ctx, filter, update)

	return err
}

func RemoveAccount(cfg config.Config, abi string, userId string, accountId model.AccountId) error {
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
	update := bson.M{"$pull": bson.M{"accounts": bson.M{"id": accountId}}}

	// Update a document
	_, err = coll.UpdateOne(ctx, filter, update)

	return err
}
