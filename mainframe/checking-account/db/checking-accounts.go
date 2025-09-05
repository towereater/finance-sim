package db

import (
	cha "mainframe-lib/checking-account/model"
	com "mainframe-lib/common/db"
	"mainframe/checking-account/config"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SelectAccount(cfg config.Config, abi string, accountId string) (cha.CheckingAccount, error) {
	// Setup timeout
	ctx, cancel := com.GetContextFromConfig(cfg.DB)
	defer cancel()

	// Retrieve the collection
	coll, err := com.GetCollection(ctx, cfg.DB, abi, cfg.Prefix, cfg.Collections.Accounts)
	if err != nil {
		return cha.CheckingAccount{}, err
	}

	// Search for a document
	var account cha.CheckingAccount
	err = coll.FindOne(ctx, bson.M{"_id": accountId}).Decode(&account)

	return account, err
}

func SelectAccountByIBAN(cfg config.Config, abi string, iban string) (cha.CheckingAccount, error) {
	// Setup timeout
	ctx, cancel := com.GetContextFromConfig(cfg.DB)
	defer cancel()

	// Retrieve the collection
	coll, err := com.GetCollection(ctx, cfg.DB, abi, cfg.Prefix, cfg.Collections.Accounts)
	if err != nil {
		return cha.CheckingAccount{}, err
	}

	// Search for a document
	var account cha.CheckingAccount
	err = coll.FindOne(ctx, bson.M{"iban": iban}).Decode(&account)

	return account, err
}

func SelectAccounts(cfg config.Config, abi string, accountFilter cha.CheckingAccount, from string, limit int) ([]cha.CheckingAccount, error) {
	// Setup timeout
	ctx, cancel := com.GetContextFromConfig(cfg.DB)
	defer cancel()

	// Retrieve the collection
	coll, err := com.GetCollection(ctx, cfg.DB, abi, cfg.Prefix, cfg.Collections.Accounts)
	if err != nil {
		return []cha.CheckingAccount{}, err
	}

	// Setup find options
	var opts options.FindOptions
	opts.SetLimit(int64(limit))

	// Setup filter
	filter := bson.M{}
	if accountFilter.IBAN != "" {
		filter["iban"] = accountFilter.IBAN
	}
	if accountFilter.Owner != "" {
		filter["owner"] = accountFilter.Owner
	}
	filter["_id"] = bson.M{"$gt": from}

	// Define the cursor
	cursor, err := coll.Find(ctx, filter, &opts)
	if err != nil {
		return []cha.CheckingAccount{}, err
	}

	// Search for the documents
	var accounts []cha.CheckingAccount
	err = cursor.All(ctx, &accounts)
	if err != nil {
		return []cha.CheckingAccount{}, err
	}

	if accounts == nil {
		return []cha.CheckingAccount{}, nil
	}
	return accounts, err
}

func InsertAccount(cfg config.Config, abi string, account cha.CheckingAccount) error {
	// Setup timeout
	ctx, cancel := com.GetContextFromConfig(cfg.DB)
	defer cancel()

	// Retrieve the collection
	coll, err := com.GetCollection(ctx, cfg.DB, abi, cfg.Prefix, cfg.Collections.Accounts)
	if err != nil {
		return err
	}

	// Insert a document
	_, err = coll.InsertOne(ctx, account)

	return err
}

func UpdateAccount(cfg config.Config, abi string, account cha.CheckingAccount) error {
	// Setup timeout
	ctx, cancel := com.GetContextFromConfig(cfg.DB)
	defer cancel()

	// Retrieve the collection
	coll, err := com.GetCollection(ctx, cfg.DB, abi, cfg.Prefix, cfg.Collections.Accounts)
	if err != nil {
		return err
	}

	// Update a document
	_, err = coll.ReplaceOne(ctx, bson.M{"_id": account.Id}, account)

	return err
}

func DeleteAccount(cfg config.Config, abi string, accountId string) error {
	// Setup timeout
	ctx, cancel := com.GetContextFromConfig(cfg.DB)
	defer cancel()

	// Retrieve the collection
	coll, err := com.GetCollection(ctx, cfg.DB, abi, cfg.Prefix, cfg.Collections.Accounts)
	if err != nil {
		return err
	}

	// Delete a document
	_, err = coll.DeleteOne(ctx, bson.M{"_id": accountId})

	return err
}
