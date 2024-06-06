package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteUser(id primitive.ObjectID) error {
	// Retrieve the collection
	coll, err := getCollection("bank", "users")
	if err != nil {
		return err
	}

	// Delete of a document
	_, err = coll.DeleteOne(context.TODO(), bson.M{"_id": id})

	return err
}
