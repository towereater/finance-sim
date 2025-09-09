package db

import (
	com "mainframe-lib/common/db"
	dos "mainframe-lib/dossier/model"
	"mainframe/dossier/config"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SelectDossier(cfg config.Config, abi string, dossierId string) (dos.Dossier, error) {
	// Setup timeout
	ctx, cancel := com.GetContextFromConfig(cfg.DB)
	defer cancel()

	// Retrieve the collection
	coll, err := com.GetCollection(ctx, cfg.DB, abi, cfg.Prefix, cfg.Collections.Dossiers)
	if err != nil {
		return dos.Dossier{}, err
	}

	// Search for a document
	var dossier dos.Dossier
	err = coll.FindOne(ctx, bson.M{"_id": dossierId}).Decode(&dossier)

	return dossier, err
}

func SelectDossiers(cfg config.Config, abi string, dossierFilter dos.Dossier, from string, limit int) ([]dos.Dossier, error) {
	// Setup timeout
	ctx, cancel := com.GetContextFromConfig(cfg.DB)
	defer cancel()

	// Retrieve the collection
	coll, err := com.GetCollection(ctx, cfg.DB, abi, cfg.Prefix, cfg.Collections.Dossiers)
	if err != nil {
		return []dos.Dossier{}, err
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
		return []dos.Dossier{}, err
	}

	// Search for the documents
	var dossiers []dos.Dossier
	err = cursor.All(ctx, &dossiers)
	if err != nil {
		return []dos.Dossier{}, err
	}

	if dossiers == nil {
		return []dos.Dossier{}, nil
	}
	return dossiers, err
}

func InsertDossier(cfg config.Config, abi string, dossier dos.Dossier) error {
	// Setup timeout
	ctx, cancel := com.GetContextFromConfig(cfg.DB)
	defer cancel()

	// Retrieve the collection
	coll, err := com.GetCollection(ctx, cfg.DB, abi, cfg.Prefix, cfg.Collections.Dossiers)
	if err != nil {
		return err
	}

	// Insert a document
	_, err = coll.InsertOne(ctx, dossier)

	return err
}

func DeleteDossier(cfg config.Config, abi string, dossierId string) error {
	// Setup timeout
	ctx, cancel := com.GetContextFromConfig(cfg.DB)
	defer cancel()

	// Retrieve the collection
	coll, err := com.GetCollection(ctx, cfg.DB, abi, cfg.Prefix, cfg.Collections.Dossiers)
	if err != nil {
		return err
	}

	// Delete a document
	_, err = coll.DeleteOne(ctx, bson.M{"_id": dossierId})

	return err
}
