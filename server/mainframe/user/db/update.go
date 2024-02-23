package db

import (
	"context"

	"mainframe/user/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateUser(id primitive.ObjectID, user model.User) error {
	// Retrieve the collection
	coll, err := getCollection("bank", "users")
	if err != nil {
		return err
	}

	// Insert of a document
	_, err = coll.UpdateOne(context.TODO(), id, user)

	return err
}
