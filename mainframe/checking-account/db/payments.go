package db

import (
	"mainframe/checking-account/config"
	"mainframe/checking-account/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SelectPayment(cfg config.Config, abi string, paymentId string) (model.Payment, error) {
	// Setup timeout
	ctx, cancel := getContextFromConfig(cfg.DB)
	defer cancel()

	// Retrieve the collection
	coll, err := getCollection(ctx, cfg.DB, abi, cfg.Prefix, cfg.Collections.Payments)
	if err != nil {
		return model.Payment{}, err
	}

	// Search for a document
	var payment model.Payment
	err = coll.FindOne(ctx, bson.M{"_id": paymentId}).Decode(&payment)

	return payment, err
}

func SelectPayments(cfg config.Config, abi string, paymentFilter model.Payment, from string, limit int) ([]model.Payment, error) {
	// Setup timeout
	ctx, cancel := getContextFromConfig(cfg.DB)
	defer cancel()

	// Retrieve the collection
	coll, err := getCollection(ctx, cfg.DB, abi, cfg.Prefix, cfg.Collections.Payments)
	if err != nil {
		return []model.Payment{}, err
	}

	// Setup find options
	var opts options.FindOptions
	opts.SetLimit(int64(limit))

	// Setup filter
	filter := bson.M{}
	if paymentFilter.Type != "" {
		filter["type"] = paymentFilter.Type
	}
	if paymentFilter.Value.Amount > 0 {
		filter["value.amount"] = paymentFilter.Value.Amount
	}
	if paymentFilter.Value.Currency != "" {
		filter["value.currency"] = paymentFilter.Value.Currency
	}
	if paymentFilter.Payer.AccountId.Account != "" {
		filter["payer.accountId.account"] = paymentFilter.Payer.AccountId.Account
	}
	if paymentFilter.Payer.AccountId.Service != "" {
		filter["payer.accountId.service"] = paymentFilter.Payer.AccountId.Service
	}
	if paymentFilter.Payee.Name != "" {
		filter["payee.name"] = paymentFilter.Payee.Name
	}
	if paymentFilter.Payee.AccountIdentification.Type != "" {
		filter["payee.accountIdentification.type"] = paymentFilter.Payee.AccountIdentification.Type
	}
	if paymentFilter.Payee.AccountIdentification.Value != "" {
		filter["payee.accountIdentification.value"] = paymentFilter.Payee.AccountIdentification.Value
	}
	if paymentFilter.Details != "" {
		filter["details"] = paymentFilter.Details
	}
	filter["_id"] = bson.M{"$gt": from}

	// Define the cursor
	cursor, err := coll.Find(ctx, filter, &opts)
	if err != nil {
		return []model.Payment{}, err
	}

	// Search for the documents
	var payments []model.Payment
	err = cursor.All(ctx, &payments)
	if err != nil {
		return []model.Payment{}, err
	}

	if payments == nil {
		return []model.Payment{}, nil
	}
	return payments, err
}

func InsertPayment(cfg config.Config, abi string, payment model.Payment) error {
	// Setup timeout
	ctx, cancel := getContextFromConfig(cfg.DB)
	defer cancel()

	// Retrieve the collection
	coll, err := getCollection(ctx, cfg.DB, abi, cfg.Prefix, cfg.Collections.Payments)
	if err != nil {
		return err
	}

	// Insert a document
	_, err = coll.InsertOne(ctx, payment)

	return err
}
