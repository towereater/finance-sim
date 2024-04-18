package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddAccount(id primitive.ObjectID, account primitive.ObjectID) error {
	// Retrieve the collection
	coll, err := getCollection("bank", "users")
	if err != nil {
		return err
	}

	// Construction of the DB objects
	filter := bson.M{"_id": id}
	update := bson.M{"$push": bson.M{"accounts": account}}

	// Insert of a document
	_, err = coll.UpdateOne(context.TODO(), filter, update)

	return err
}
