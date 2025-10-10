package db

import (
	cha "mainframe-lib/checking-account/model"
	dcom "mainframe-lib/common/db"
	scom "mainframe-lib/common/service"
	"processor/payment/config"

	"go.mongodb.org/mongo-driver/bson"
)

func SelectAccount(db config.DB, abi string, accountId string) (cha.CheckingAccount, error) {
	// Setup timeout
	ctx, cancel := scom.GetContextWithTimeout(db.Timeout)
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
	ctx, cancel := scom.GetContextWithTimeout(db.Timeout)
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

func UpdateAccount(db config.DB, abi string, account cha.CheckingAccount) error {
	// Setup timeout
	ctx, cancel := scom.GetContextWithTimeout(db.Timeout)
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
