package db

import (
	"context"

	"mainframe/account/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SelectAccount(id primitive.ObjectID) (*model.Account, error) {
	// Retrieve the collection
	coll, err := getCollection("bank", "accounts")
	if err != nil {
		return nil, err
	}

	// Search for a document
	var user model.Account
	err = coll.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&user)

	return &user, err
}

func SelectAccounts(account model.Account, fromId primitive.ObjectID, limit int, order int) ([]model.Account, error) {
	// Retrieve the collection
	coll, err := getCollection("bank", "accounts")
	if err != nil {
		return nil, err
	}

	// Setting up find options
	var opts options.FindOptions
	opts.SetLimit(int64(limit))
	opts.SetSort(bson.D{bson.E{Key: "_id", Value: order}})

	// Setting up filter
	var filter = make(bson.D, 0, 6)
	if account.IBAN != "" {
		filter = append(filter, bson.E{Key: "iban", Value: account.IBAN})
	}
	if account.Owner != "" {
		filter = append(filter, bson.E{Key: "owner", Value: account.Owner})
	}

	if fromId != primitive.NilObjectID {
		if order == 1 {
			filter = append(filter, bson.E{Key: "$gt", Value: fromId})
		} else {
			filter = append(filter, bson.E{Key: "$lt", Value: fromId})
		}
	}

	// Definition of the cursor
	cursor, err := coll.Find(context.TODO(), filter, &opts)
	if err != nil {
		return nil, err
	}

	// Search for the documents
	var accounts []model.Account
	err = cursor.All(context.TODO(), &accounts)

	return accounts, err
}
