package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteAccount(id primitive.ObjectID) error {
	// Retrieve the collection
	coll, err := getCollection("bank", "accounts")
	if err != nil {
		return err
	}

	// Delete of a document
	_, err = coll.DeleteOne(context.TODO(), bson.M{"_id": id})

	return err
}
