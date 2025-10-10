package service

import (
	"context"
	"fmt"
	cha "mainframe-lib/checking-account/model"
	qcom "mainframe-lib/common/queue"
	"processor/payment/config"
	"processor/payment/db"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

func ProcessPayment(cfg config.Config, abi string, payment cha.Payment) (cha.PaymentOutcome, error) {
	// Return variables
	outcome := cha.PaymentOutcome{
		Timestamp: time.Now().Format(time.DateTime),
	}
	var err error

	// Obtain payer account
	var payerAccount cha.CheckingAccount

	switch payment.Payer.AccountIdentification.Type {
	case "ID":
		payerAccount, err = db.SelectAccount(cfg.DB, abi, payment.Payer.AccountIdentification.Value)
	case "IBAN":
		payerAccount, err = db.SelectAccountByIBAN(cfg.DB, abi, payment.Payer.AccountIdentification.Value)
	}
	if err == mongo.ErrNoDocuments {
		err = fmt.Errorf("payer account with %s %s not found",
			payment.Payer.AccountIdentification.Type,
			payment.Payer.AccountIdentification.Value)
		outcome.Status = "E"
		outcome.Message = err.Error()

		return outcome, err
	}
	if err != nil {
		outcome.Status = "E"
		outcome.Message = err.Error()

		return outcome, err
	}

	// Check payer account cash availability
	if payerAccount.Value.Amount < payment.Value.Amount {
		err = fmt.Errorf("payer account %s without enough funds", payerAccount.Id)
		outcome.Status = "E"
		outcome.Message = err.Error()

		return outcome, err
	}

	// Obtain payee account
	var payeeAccount cha.CheckingAccount

	if payment.Payee.AccountIdentification.Type == "IBAN" {
		payeeAccount, err = db.SelectAccountByIBAN(cfg.DB, abi, payment.Payee.AccountIdentification.Value)
		if err == mongo.ErrNoDocuments {
			err = fmt.Errorf("payee account with %s %s not found",
				payment.Payee.AccountIdentification.Type,
				payment.Payee.AccountIdentification.Value)
			outcome.Status = "E"
			outcome.Message = err.Error()

			return outcome, err
		}
		if err != nil {
			outcome.Status = "E"
			outcome.Message = err.Error()

			return outcome, err
		}
	} else {
		err = fmt.Errorf("unsupported payee account identification type %s",
			payment.Payee.AccountIdentification.Type)
		outcome.Status = "E"
		outcome.Message = err.Error()

		return outcome, err
	}

	// Remove cash from payer account
	payerAccount.Value.Amount -= payment.Value.Amount
	err = db.UpdateAccount(cfg.DB, abi, payerAccount)
	if err != nil {
		outcome.Status = "E"
		outcome.Message = err.Error()

		return outcome, err
	}

	// Add cash to payee account
	payeeAccount.Value.Amount += payment.Value.Amount
	err = db.UpdateAccount(cfg.DB, abi, payeeAccount)
	if err != nil {
		outcome.Status = "E"
		outcome.Message = err.Error()

		return outcome, err
	}

	outcome.Status = "O"
	outcome.Message = "payment processed correctly"

	return outcome, err
}

func UnqueuePayment(queue config.Queue) (string, string, error) {
	// Setup timeout
	ctx := context.Background()

	// Unqueue document
	key, value, err := qcom.UnqueueContent(ctx, queue.Queue, queue.Topics.Payments)

	return key, value, err
}
