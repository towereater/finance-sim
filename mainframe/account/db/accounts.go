package db

import (
	"mainframe/account/config"
	"mainframe/account/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SelectAccount(cfg config.Config, abi string, accountId primitive.ObjectID) (model.Account, error) {
	// Setup timeout
	ctx, cancel := getContextFromConfig(cfg.DB)
	defer cancel()

	// Retrieve the collection
	coll, err := getCollection(ctx, cfg.DB, abi, cfg.Prefix, cfg.Collections.Accounts)
	if err != nil {
		return model.Account{}, err
	}

	// Search for a document
	var account model.Account
	err = coll.FindOne(ctx, bson.M{"_id": accountId}).Decode(&account)

	return account, err
}

func SelectAccounts(cfg config.Config, abi string, accountFilter model.Account, from primitive.ObjectID, limit int) ([]model.Account, error) {
	// Setup timeout
	ctx, cancel := getContextFromConfig(cfg.DB)
	defer cancel()

	// Retrieve the collection
	coll, err := getCollection(ctx, cfg.DB, abi, cfg.Prefix, cfg.Collections.Accounts)
	if err != nil {
		return []model.Account{}, err
	}

	// Setup find options
	var opts options.FindOptions
	opts.SetLimit(int64(limit))

	// Setup filter
	filter := bson.M{}
	if accountFilter.Owner != "" {
		filter["owner"] = accountFilter.Owner
	}
	if accountFilter.Service != "" {
		filter["service"] = accountFilter.Service
	}
	filter["_id"] = bson.M{"$gt": from}

	// Define the cursor
	cursor, err := coll.Find(ctx, filter, &opts)
	if err != nil {
		return []model.Account{}, err
	}

	// Search for the documents
	var accounts []model.Account
	err = cursor.All(ctx, &accounts)
	if err != nil {
		return []model.Account{}, err
	}

	if accounts == nil {
		return []model.Account{}, nil
	}
	return accounts, err
}

func InsertAccount(cfg config.Config, abi string, account model.Account) error {
	// Setup timeout
	ctx, cancel := getContextFromConfig(cfg.DB)
	defer cancel()

	// Retrieve the collection
	coll, err := getCollection(ctx, cfg.DB, abi, cfg.Prefix, cfg.Collections.Accounts)
	if err != nil {
		return err
	}

	// Insert a document
	_, err = coll.InsertOne(ctx, account)

	return err
}

func DeleteAccount(cfg config.Config, abi string, accountId primitive.ObjectID) error {
	// Setup timeout
	ctx, cancel := getContextFromConfig(cfg.DB)
	defer cancel()

	// Retrieve the collection
	coll, err := getCollection(ctx, cfg.DB, abi, cfg.Prefix, cfg.Collections.Accounts)
	if err != nil {
		return err
	}

	// Delete a document
	_, err = coll.DeleteOne(ctx, bson.M{"_id": accountId})

	return err
}
