package db

import (
	"context"

	"mainframe/user/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateUser(id primitive.ObjectID, user model.User) error {
	// Retrieve the collection
	coll, err := getCollection("bank", "users")
	if err != nil {
		return err
	}

	// Construction of the DB objects
	filter := bson.M{"_id": id}
	update := bson.M{"$set": user}

	// Insert of a document
	_, err = coll.UpdateOne(context.TODO(), filter, update)

	return err
}
