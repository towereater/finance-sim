package db

import (
	cha "mainframe-lib/checking-account/model"
	dcom "mainframe-lib/common/db"
	scom "mainframe-lib/common/service"
	"processor/payment/config"

	"go.mongodb.org/mongo-driver/bson"
)

func SelectPayment(db config.DB, abi string, paymentId string) (cha.Payment, error) {
	// Setup timeout
	ctx, cancel := scom.GetContextWithTimeout(db.Timeout)
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

func UpdatePayment(db config.DB, abi string, payment cha.Payment) error {
	// Setup timeout
	ctx, cancel := scom.GetContextWithTimeout(db.Timeout)
	defer cancel()

	// Retrieve the collection
	coll, err := dcom.GetCollection(ctx, db.DB, abi, db.Collections.Payments)
	if err != nil {
		return err
	}

	// Update a document
	_, err = coll.ReplaceOne(ctx, bson.M{"_id": payment.Id}, payment)

	return err
}
