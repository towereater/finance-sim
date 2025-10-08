package db

import (
	cha "mainframe-lib/checking-account/model"
	dcom "mainframe-lib/common/db"
	scom "mainframe-lib/common/service"
	"mainframe/checking-account/config"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SelectAccount(db config.DB, abi string, accountId string) (cha.CheckingAccount, error) {
	// Setup timeout
	ctx, cancel := scom.GetContextWithTimeout(db.DB.Timeout)
	defer cancel()

	// Retrieve the collection
	coll, err := dcom.GetCollection(ctx, db.DB, abi, db.Collections.Accounts)
	if err != nil {
		return cha.CheckingAccount{}, err
	}

	// Search for a document
	var account cha.CheckingAccount
	err = coll.FindOne(ctx, bson.M{"_id": accountId}).Decode(&account)

	return account, err
}

func SelectAccountByIBAN(db config.DB, abi string, iban string) (cha.CheckingAccount, error) {
	// Setup timeout
	ctx, cancel := scom.GetContextWithTimeout(db.DB.Timeout)
	defer cancel()

	// Retrieve the collection
	coll, err := dcom.GetCollection(ctx, db.DB, abi, db.Collections.Accounts)
	if err != nil {
		return cha.CheckingAccount{}, err
	}

	// Search for a document
	var account cha.CheckingAccount
	err = coll.FindOne(ctx, bson.M{"iban": iban}).Decode(&account)

	return account, err
}

func SelectAccounts(db config.DB, abi string, accountFilter cha.CheckingAccount, from string, limit int) ([]cha.CheckingAccount, error) {
	// Setup timeout
	ctx, cancel := scom.GetContextWithTimeout(db.DB.Timeout)
	defer cancel()

	// Retrieve the collection
	coll, err := dcom.GetCollection(ctx, db.DB, abi, db.Collections.Accounts)
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

func InsertAccount(db config.DB, abi string, account cha.CheckingAccount) error {
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

func UpdateAccount(db config.DB, abi string, account cha.CheckingAccount) error {
	// Setup timeout
	ctx, cancel := scom.GetContextWithTimeout(db.DB.Timeout)
	defer cancel()

	// Retrieve the collection
	coll, err := dcom.GetCollection(ctx, db.DB, abi, db.Collections.Accounts)
	if err != nil {
		return err
	}

	// Update a document
	_, err = coll.ReplaceOne(ctx, bson.M{"_id": account.Id}, account)

	return err
}

func DeleteAccount(db config.DB, abi string, accountId string) error {
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
