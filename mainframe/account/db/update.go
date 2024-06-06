package db

import (
	"context"

	"mainframe/account/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateAccount(id primitive.ObjectID, account model.Account) error {
	// Retrieve the collection
	coll, err := getCollection("bank", "accounts")
	if err != nil {
		return err
	}

	// Construction of the DB objects
	filter := bson.M{"_id": id}
	update := bson.M{"$set": account}

	// Insert of a document
	_, err = coll.UpdateOne(context.TODO(), filter, update)

	return err
}
