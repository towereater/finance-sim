package db

import (
	"context"

	"mainframe/user/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SelectUser(id primitive.ObjectID) (*model.User, error) {
	// Retrieve the collection
	coll, err := getCollection("bank", "users")
	if err != nil {
		return nil, err
	}

	// Search for a document
	var user model.User
	err = coll.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&user)

	return &user, err
}

func SelectUsers(user model.User, fromId primitive.ObjectID, limit int, order int) ([]model.User, error) {
	// Retrieve the collection
	coll, err := getCollection("bank", "users")
	if err != nil {
		return nil, err
	}

	// Setting up find options
	var opts options.FindOptions
	opts.SetLimit(int64(limit))
	opts.SetSort(bson.D{bson.E{Key: "_id", Value: order}})

	// Setting up filter
	var filter = make(bson.D, 0, 6)
	if user.Username != "" {
		filter = append(filter, bson.E{Key: "username", Value: user.Username})
	}
	if user.Password != "" {
		filter = append(filter, bson.E{Key: "password", Value: user.Password})
	}
	if user.Name != "" {
		filter = append(filter, bson.E{Key: "name", Value: user.Name})
	}
	if user.Surname != "" {
		filter = append(filter, bson.E{Key: "surname", Value: user.Surname})
	}
	if user.Birth != "" {
		filter = append(filter, bson.E{Key: "birth", Value: user.Birth})
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
	var users []model.User
	err = cursor.All(context.TODO(), &users)

	return users, err
}
