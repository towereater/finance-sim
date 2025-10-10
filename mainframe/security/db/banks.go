package db

import (
	dcom "mainframe-lib/common/db"
	scom "mainframe-lib/common/service"
	sec "mainframe-lib/security/model"
	"mainframe/security/config"

	"go.mongodb.org/mongo-driver/bson"
)

func SelectBankByAbi(db config.DB, abi string) (sec.Bank, error) {
	// Setup timeout
	ctx, cancel := scom.GetContextWithTimeout(db.Timeout)
	defer cancel()

	// Retrieve the collection
	coll, err := dcom.GetCollection(ctx, db.DB, "09999", db.Collections.Banks)
	if err != nil {
		return sec.Bank{}, err
	}

	// Search for a document
	var bank sec.Bank
	err = coll.FindOne(ctx, bson.M{"_id": abi}).Decode(&bank)

	return bank, err
}

func InsertBank(db config.DB, bank sec.Bank) error {
	// Setup timeout
	ctx, cancel := scom.GetContextWithTimeout(db.Timeout)
	defer cancel()

	// Retrieve the collection
	coll, err := dcom.GetCollection(ctx, db.DB, "09999", db.Collections.Banks)
	if err != nil {
		return err
	}

	// Insert a document
	_, err = coll.InsertOne(ctx, bank)

	return err
}

func DeleteBank(db config.DB, abi string) error {
	// Setup timeout
	ctx, cancel := scom.GetContextWithTimeout(db.Timeout)
	defer cancel()

	// Retrieve the collection
	coll, err := dcom.GetCollection(ctx, db.DB, "09999", db.Collections.Banks)
	if err != nil {
		return err
	}

	// Delete a document
	_, err = coll.DeleteOne(ctx, bson.M{"_id": abi})

	return err
}
