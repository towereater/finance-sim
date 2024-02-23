package db

import (
	"context"

	"mainframe/user/model"
)

func InsertUser(user model.User) error {
	// Retrieve the collection
	coll, err := getCollection("bank", "users")
	if err != nil {
		return err
	}

	// Insert of a document
	_, err = coll.InsertOne(context.TODO(), user)

	return err
}
