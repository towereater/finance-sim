package db

import (
	dcom "mainframe-lib/common/db"
	scom "mainframe-lib/common/service"
	sec "mainframe-lib/security/model"
	"mainframe/security/config"

	"go.mongodb.org/mongo-driver/bson"
)

func SelectUserByApiKey(db config.DB, abi string, apiKey string) (sec.User, error) {
	// Setup timeout
	ctx, cancel := scom.GetContextWithTimeout(db.Timeout)
	defer cancel()

	// Retrieve the collection
	coll, err := dcom.GetCollection(ctx, db.DB, abi, db.Collections.Users)
	if err != nil {
		return sec.User{}, err
	}

	// Search for a document
	var user sec.User
	err = coll.FindOne(ctx, bson.M{"apiKey": apiKey}).Decode(&user)

	return user, err
}

func InsertUser(db config.DB, abi string, user sec.User) error {
	// Setup timeout
	ctx, cancel := scom.GetContextWithTimeout(db.Timeout)
	defer cancel()

	// Retrieve the collection
	coll, err := dcom.GetCollection(ctx, db.DB, abi, db.Collections.Users)
	if err != nil {
		return err
	}

	// Insert a document
	_, err = coll.InsertOne(ctx, user)

	return err
}

func DeleteUser(db config.DB, abi string, userId string) error {
	// Setup timeout
	ctx, cancel := scom.GetContextWithTimeout(db.Timeout)
	defer cancel()

	// Retrieve the collection
	coll, err := dcom.GetCollection(ctx, db.DB, abi, db.Collections.Users)
	if err != nil {
		return err
	}

	// Delete a document
	_, err = coll.DeleteOne(ctx, bson.M{"_id": userId})

	return err
}
