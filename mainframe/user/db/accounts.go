package db

import (
	dcom "mainframe-lib/common/db"
	scom "mainframe-lib/common/service"
	usr "mainframe-lib/user/model"
	"mainframe/user/config"

	"go.mongodb.org/mongo-driver/bson"
)

func AddAccount(db config.DB, abi string, userId string, account usr.Account) error {
	// Setup timeout
	ctx, cancel := scom.GetContextWithTimeout(db.DB.Timeout)
	defer cancel()

	// Retrieve the collection
	coll, err := dcom.GetCollection(ctx, db.DB, abi, db.Collections.Users)
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

func RemoveAccount(db config.DB, abi string, userId string, accountId usr.AccountId) error {
	// Setup timeout
	ctx, cancel := scom.GetContextWithTimeout(db.DB.Timeout)
	defer cancel()

	// Retrieve the collection
	coll, err := dcom.GetCollection(ctx, db.DB, abi, db.Collections.Users)
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
