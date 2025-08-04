package db

import (
	"mainframe/security/config"
	"mainframe/security/model"

	"go.mongodb.org/mongo-driver/bson"
)

func SelectApiKey(cfg config.Config, abi string, apiKeyId string) (model.ApiKey, error) {
	// Setup timeout
	ctx, cancel := getContextFromConfig(cfg.DB)
	defer cancel()

	// Retrieve the collection
	coll, err := getCollection(ctx, cfg.DB, abi, cfg.Prefix, cfg.Collections.ApiKeys)
	if err != nil {
		return model.ApiKey{}, err
	}

	// Search for a document
	var apiKey model.ApiKey
	err = coll.FindOne(ctx, bson.M{"_id": apiKeyId}).Decode(&apiKey)

	return apiKey, err
}

func InsertApiKey(cfg config.Config, abi string, apiKey model.ApiKey) error {
	// Setup timeout
	ctx, cancel := getContextFromConfig(cfg.DB)
	defer cancel()

	// Retrieve the collection
	coll, err := getCollection(ctx, cfg.DB, abi, cfg.Prefix, cfg.Collections.ApiKeys)
	if err != nil {
		return err
	}

	// Insert a document
	_, err = coll.InsertOne(ctx, apiKey)

	return err
}

func DeleteApiKey(cfg config.Config, abi string, apiKeyId string) error {
	// Setup timeout
	ctx, cancel := getContextFromConfig(cfg.DB)
	defer cancel()

	// Retrieve the collection
	coll, err := getCollection(ctx, cfg.DB, abi, cfg.Prefix, cfg.Collections.ApiKeys)
	if err != nil {
		return err
	}

	// Delete a document
	_, err = coll.DeleteOne(ctx, bson.M{"_id": apiKeyId})

	return err
}
