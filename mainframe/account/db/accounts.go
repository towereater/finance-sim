package db

import (
	acc "mainframe-lib/account/model"
	dcom "mainframe-lib/common/db"
	scom "mainframe-lib/common/service"
	"mainframe/account/config"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SelectAccount(db config.DB, abi string, accountId acc.AccountId) (acc.Account, error) {
	// Setup timeout
	ctx, cancel := scom.GetContextWithTimeout(db.DB.Timeout)
	defer cancel()

	// Retrieve the collection
	coll, err := dcom.GetCollection(ctx, db.DB, abi, db.Collections.Accounts)
	if err != nil {
		return acc.Account{}, err
	}

	// Search for a document
	var account acc.Account
	err = coll.FindOne(ctx, bson.M{"_id": accountId}).Decode(&account)

	return account, err
}

func SelectAccounts(db config.DB, abi string, accountFilter acc.Account, from string, limit int) ([]acc.Account, error) {
	// Setup timeout
	ctx, cancel := scom.GetContextWithTimeout(db.DB.Timeout)
	defer cancel()

	// Retrieve the collection
	coll, err := dcom.GetCollection(ctx, db.DB, abi, db.Collections.Accounts)
	if err != nil {
		return []acc.Account{}, err
	}

	// Setup find options
	var opts options.FindOptions
	opts.SetLimit(int64(limit))

	// Setup filter
	filter := bson.M{}
	if accountFilter.Id.Account != "" {
		filter["_id.account"] = accountFilter.Id.Account
	}
	if accountFilter.Id.Service != "" {
		filter["_id.service"] = accountFilter.Id.Service
	}
	if accountFilter.Owner != "" {
		filter["owner"] = accountFilter.Owner
	}
	filter["_id.account"] = bson.M{"$gt": from}

	// Define the cursor
	cursor, err := coll.Find(ctx, filter, &opts)
	if err != nil {
		return []acc.Account{}, err
	}

	// Search for the documents
	var accounts []acc.Account
	err = cursor.All(ctx, &accounts)
	if err != nil {
		return []acc.Account{}, err
	}

	if accounts == nil {
		return []acc.Account{}, nil
	}
	return accounts, err
}

func InsertAccount(db config.DB, abi string, account acc.Account) error {
	// Setup timeout
	ctx, cancel := scom.GetContextWithTimeout(db.DB.Timeout)
	defer cancel()

	// Retrieve the collection
	coll, err := dcom.GetCollection(ctx, db.DB, abi, db.Collections.Accounts)
	if err != nil {
		return err
	}

	// Insert a document
	_, err = coll.InsertOne(ctx, account)

	return err
}

func DeleteAccount(db config.DB, abi string, accountId acc.AccountId) error {
	// Setup timeout
	ctx, cancel := scom.GetContextWithTimeout(db.DB.Timeout)
	defer cancel()

	// Retrieve the collection
	coll, err := dcom.GetCollection(ctx, db.DB, abi, db.Collections.Accounts)
	if err != nil {
		return err
	}

	// Delete a document
	_, err = coll.DeleteOne(ctx, bson.M{"_id": accountId})

	return err
}
