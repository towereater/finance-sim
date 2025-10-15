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
	if payment.Payer.Account.Id == "" {
		err = fmt.Errorf("payer account identification missing")
		outcome.Status = "E"
		outcome.Message = err.Error()

		return outcome, err
	}

	payerAccount, err := db.SelectAccount(cfg.DB, abi, payment.Payer.Account.Id)
	if err == mongo.ErrNoDocuments {
		err = fmt.Errorf("payer account %s not found", payment.Payee.Account.Id)
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
	if payment.Payee.Account.Id == "" {
		err = fmt.Errorf("payee account identification missing")
		outcome.Status = "E"
		outcome.Message = err.Error()

		return outcome, err
	}

	payeeAccount, err := db.SelectAccount(cfg.DB, abi, payment.Payee.Account.Id)
	if err == mongo.ErrNoDocuments {
		err = fmt.Errorf("payee account %s not found", payment.Payee.Account.Id)
		outcome.Status = "E"
		outcome.Message = err.Error()

		return outcome, err
	}
	if err != nil {
		outcome.Status = "E"
		outcome.Message = err.Error()

		return outcome, err
	}

	// Set up of the successful outcome status
	outcome.Status = "O"
	outcome.Message = "payment processed correctly"

	// Construction of checking account payment
	ckPayment := cha.CheckingPayment{
		Id:      payment.Id,
		Type:    payment.Type,
		Value:   payment.Value,
		Payee:   payment.Payee,
		Details: payment.Details,
		Outcome: outcome,
	}

	// Update payer account
	payerAccount.Value.Amount -= payment.Value.Amount
	payerAccount.LastPayments = append(payerAccount.LastPayments, ckPayment)
	if len(payerAccount.LastPayments) > 5 {
		payerAccount.LastPayments = payerAccount.LastPayments[len(payerAccount.LastPayments)-5:]
	}

	err = db.UpdateAccount(cfg.DB, abi, payerAccount)
	if err != nil {
		outcome.Status = "E"
		outcome.Message = err.Error()

		return outcome, err
	}

	// Update to payee account
	payeeAccount.Value.Amount += payment.Value.Amount
	payeeAccount.LastPayments = append(payeeAccount.LastPayments, ckPayment)
	if len(payeeAccount.LastPayments) > 5 {
		payeeAccount.LastPayments = payeeAccount.LastPayments[len(payeeAccount.LastPayments)-5:]
	}

	err = db.UpdateAccount(cfg.DB, abi, payeeAccount)
	if err != nil {
		outcome.Status = "E"
		outcome.Message = err.Error()

		return outcome, err
	}

	return outcome, err
}

func UnqueuePayment(queue config.Queue) (string, string, error) {
	// Setup timeout
	ctx := context.Background()

	// Unqueue document
	key, value, err := qcom.UnqueueContent(ctx, queue.Queue, queue.Topics.Payments)

	return key, value, err
}
