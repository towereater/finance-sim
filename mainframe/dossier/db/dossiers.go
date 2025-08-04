package db

import (
	"mainframe/dossier/config"
	"mainframe/dossier/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SelectDossier(cfg config.Config, abi string, dossierId string) (model.Dossier, error) {
	// Setup timeout
	ctx, cancel := getContextFromConfig(cfg.DB)
	defer cancel()

	// Retrieve the collection
	coll, err := getCollection(ctx, cfg.DB, abi, cfg.Prefix, cfg.Collections.Dossiers)
	if err != nil {
		return model.Dossier{}, err
	}

	// Search for a document
	var dossier model.Dossier
	err = coll.FindOne(ctx, bson.M{"_id": dossierId}).Decode(&dossier)

	return dossier, err
}

func SelectDossiers(cfg config.Config, abi string, dossierFilter model.Dossier, from string, limit int) ([]model.Dossier, error) {
	// Setup timeout
	ctx, cancel := getContextFromConfig(cfg.DB)
	defer cancel()

	// Retrieve the collection
	coll, err := getCollection(ctx, cfg.DB, abi, cfg.Prefix, cfg.Collections.Dossiers)
	if err != nil {
		return []model.Dossier{}, err
	}

	// Setup find options
	var opts options.FindOptions
	opts.SetLimit(int64(limit))

	// Setup filter
	filter := bson.M{}
	if dossierFilter.Owner != "" {
		filter["owner"] = dossierFilter.Owner
	}
	filter["_id"] = bson.M{"$gt": from}

	// Define the cursor
	cursor, err := coll.Find(ctx, filter, &opts)
	if err != nil {
		return []model.Dossier{}, err
	}

	// Search for the documents
	var dossiers []model.Dossier
	err = cursor.All(ctx, &dossiers)
	if err != nil {
		return []model.Dossier{}, err
	}

	if dossiers == nil {
		return []model.Dossier{}, nil
	}
	return dossiers, err
}

func InsertDossier(cfg config.Config, abi string, dossier model.Dossier) error {
	// Setup timeout
	ctx, cancel := getContextFromConfig(cfg.DB)
	defer cancel()

	// Retrieve the collection
	coll, err := getCollection(ctx, cfg.DB, abi, cfg.Prefix, cfg.Collections.Dossiers)
	if err != nil {
		return err
	}

	// Insert a document
	_, err = coll.InsertOne(ctx, dossier)

	return err
}

func DeleteDossier(cfg config.Config, abi string, dossierId string) error {
	// Setup timeout
	ctx, cancel := getContextFromConfig(cfg.DB)
	defer cancel()

	// Retrieve the collection
	coll, err := getCollection(ctx, cfg.DB, abi, cfg.Prefix, cfg.Collections.Dossiers)
	if err != nil {
		return err
	}

	// Delete a document
	_, err = coll.DeleteOne(ctx, bson.M{"_id": dossierId})

	return err
}
