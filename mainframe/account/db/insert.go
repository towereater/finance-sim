package db

import (
	"context"

	"mainframe/account/model"
)

func InsertAccount(account model.Account) error {
	// Retrieve the collection
	coll, err := getCollection("bank", "accounts")
	if err != nil {
		return err
	}

	// Insert of a document
	_, err = coll.InsertOne(context.TODO(), account)

	return err
}
