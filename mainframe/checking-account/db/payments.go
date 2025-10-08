package db

import (
	cha "mainframe-lib/checking-account/model"
	dcom "mainframe-lib/common/db"
	scom "mainframe-lib/common/service"
	"mainframe/checking-account/config"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SelectPayment(db config.DB, abi string, paymentId string) (cha.Payment, error) {
	// Setup timeout
	ctx, cancel := scom.GetContextWithTimeout(db.DB.Timeout)
	defer cancel()

	// Retrieve the collection
	coll, err := dcom.GetCollection(ctx, db.DB, abi, db.Collections.Payments)
	if err != nil {
		return cha.Payment{}, err
	}

	// Search for a document
	var payment cha.Payment
	err = coll.FindOne(ctx, bson.M{"_id": paymentId}).Decode(&payment)

	return payment, err
}

func SelectPayments(db config.DB, abi string, paymentFilter cha.Payment, from string, limit int) ([]cha.Payment, error) {
	// Setup timeout
	ctx, cancel := scom.GetContextWithTimeout(db.DB.Timeout)
	defer cancel()

	// Retrieve the collection
	coll, err := dcom.GetCollection(ctx, db.DB, abi, db.Collections.Payments)
	if err != nil {
		return []cha.Payment{}, err
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
	if paymentFilter.Payer.AccountIdentification.Type != "" {
		filter["payer.accountIdentification.type"] = paymentFilter.Payer.AccountIdentification.Type
	}
	if paymentFilter.Payer.AccountIdentification.Value != "" {
		filter["payer.accountIdentification.value"] = paymentFilter.Payer.AccountIdentification.Value
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
		return []cha.Payment{}, err
	}

	// Search for the documents
	var payments []cha.Payment
	err = cursor.All(ctx, &payments)
	if err != nil {
		return []cha.Payment{}, err
	}

	if payments == nil {
		return []cha.Payment{}, nil
	}
	return payments, err
}

func InsertPayment(db config.DB, abi string, payment cha.Payment) error {
	// Setup timeout
	ctx, cancel := scom.GetContextWithTimeout(db.DB.Timeout)
	defer cancel()

	// Retrieve the collection
	coll, err := dcom.GetCollection(ctx, db.DB, abi, db.Collections.Payments)
	if err != nil {
		return err
	}

	// Insert a document
	_, err = coll.InsertOne(ctx, payment)

	return err
}
