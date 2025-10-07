package db

import (
	com "mainframe-lib/common/db"
	sec "mainframe-lib/security/model"
	"mainframe/security/config"

	"go.mongodb.org/mongo-driver/bson"
)

func SelectBankByAbi(cfg config.DBConfig, abi string) (sec.Bank, error) {
	// Setup timeout
	ctx, cancel := com.GetContextFromConfig(cfg.DBConfig)
	defer cancel()

	// Retrieve the collection
	coll, err := com.GetCollection(ctx, cfg.DBConfig, "09999", cfg.Collections.Banks)
	if err != nil {
		return sec.Bank{}, err
	}

	// Search for a document
	var bank sec.Bank
	err = coll.FindOne(ctx, bson.M{"_id": abi}).Decode(&bank)

	return bank, err
}

func InsertBank(cfg config.DBConfig, bank sec.Bank) error {
	// Setup timeout
	ctx, cancel := com.GetContextFromConfig(cfg.DBConfig)
	defer cancel()

	// Retrieve the collection
	coll, err := com.GetCollection(ctx, cfg.DBConfig, "09999", cfg.Collections.Banks)
	if err != nil {
		return err
	}

	// Insert a document
	_, err = coll.InsertOne(ctx, bank)

	return err
}

func DeleteBank(cfg config.DBConfig, abi string) error {
	// Setup timeout
	ctx, cancel := com.GetContextFromConfig(cfg.DBConfig)
	defer cancel()

	// Retrieve the collection
	coll, err := com.GetCollection(ctx, cfg.DBConfig, "09999", cfg.Collections.Banks)
	if err != nil {
		return err
	}

	// Delete a document
	_, err = coll.DeleteOne(ctx, bson.M{"_id": abi})

	return err
}
