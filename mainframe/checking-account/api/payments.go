package api

import (
	"encoding/json"
	"fmt"
	"mainframe/checking-account/config"
	"mainframe/checking-account/db"
	"mainframe/checking-account/service"
	"net/http"
	"strconv"

	cha "mainframe-lib/checking-account/model"
	com "mainframe-lib/common/config"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetPayment(w http.ResponseWriter, r *http.Request) {
	// Extract path parameters
	paymentId := r.PathValue(string(config.ContextPaymentId))
	if len(paymentId) != 24 {
		fmt.Printf("Invalid payment id value\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Extract context parameters
	cfg := r.Context().Value(com.ContextConfig).(config.Config)
	abi := r.Context().Value(com.ContextAbi).(string)

	// Select the document
	payment, err := db.SelectPayment(cfg.DB, abi, paymentId)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No payments with id %s\n", paymentId)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		fmt.Printf("Error while searching payment with id %s: %s\n", paymentId, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Response output
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(payment)
}

func GetPayments(w http.ResponseWriter, r *http.Request) {
	// Extract query parameters
	queryParams := r.URL.Query()
	var err error

	from := queryParams.Get(string(config.ContextFrom))
	if from != "" && len(from) != 24 {
		fmt.Printf("Invalid %s parameter\n", string(config.ContextFrom))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	limit := 50
	if queryParams.Has(string(config.ContextLimit)) {
		limit, err = strconv.Atoi(queryParams.Get(string(config.ContextLimit)))

		if err != nil {
			fmt.Printf("Invalid %s parameter\n", string(config.ContextLimit))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if limit > 50 {
			limit = 50
		}
	}

	// Build the filter
	var filter cha.Payment
	filter.Type = queryParams.Get("paymenyType")
	amount, err := strconv.ParseFloat(queryParams.Get("amount"), 32)
	if err != nil {
		filter.Value.Amount = float32(amount)
	}
	filter.Value.Currency = queryParams.Get("currency")
	filter.Payer.AccountIdentification.Type = queryParams.Get("payerType")
	filter.Payer.AccountIdentification.Value = queryParams.Get("payerValue")
	filter.Payee.Name = queryParams.Get("name")
	filter.Payee.AccountIdentification.Type = queryParams.Get("payeeType")
	filter.Payee.AccountIdentification.Value = queryParams.Get("payeeValue")
	filter.Details = queryParams.Get("details")

	// Extract context parameters
	cfg := r.Context().Value(com.ContextConfig).(config.Config)
	abi := r.Context().Value(com.ContextAbi).(string)

	// Select all documents
	payments, err := db.SelectPayments(cfg.DB, abi, filter, from, limit)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No payments with filter %+v\n", filter)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		fmt.Printf("Error while searching payments with filter %+v: %s\n", filter, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Response output
	if len(payments) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(payments)
}

func InsertPayment(w http.ResponseWriter, r *http.Request) {
	// Parse the request
	var req cha.InsertPayment
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		fmt.Printf("Could not convert request body\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if req.Type != "BANK_TRANSFER" {
		fmt.Printf("Invalid payment type\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if req.Value.Amount <= 0 {
		fmt.Printf("Invalid value amount\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if len(req.Value.Currency) != 3 {
		fmt.Printf("Invalid value currency\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if req.Payer.AccountIdentification.Type != "ID" &&
		req.Payer.AccountIdentification.Type != "IBAN" {
		fmt.Printf("Invalid payer account identification type\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if len(req.Payer.AccountIdentification.Value) != 24 &&
		len(req.Payer.AccountIdentification.Value) != 27 {
		fmt.Printf("Invalid payer account value\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if req.Payee.Name == "" {
		fmt.Printf("Invalid payee name\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if req.Payee.AccountIdentification.Type != "IBAN" {
		fmt.Printf("Invalid payee account identification type\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if len(req.Payee.AccountIdentification.Value) != 27 {
		fmt.Printf("Invalid payee account identification value\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Extract context parameters
	cfg := r.Context().Value(com.ContextConfig).(config.Config)
	abi := r.Context().Value(com.ContextAbi).(string)

	// Obtain payer account and check its cash availability
	var payerAccount cha.CheckingAccount

	switch req.Payer.AccountIdentification.Type {
	case "ID":
		payerAccount, err = db.SelectAccount(cfg.DB, abi, req.Payer.AccountIdentification.Value)
		if err != nil {
			fmt.Printf("Error while searching payer account with id %s: %s\n",
				req.Payer.AccountIdentification.Value, err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if payerAccount.Id == "" {
			fmt.Printf("Payer account with id %s not found\n", req.Payer.AccountIdentification.Value)
			w.WriteHeader(http.StatusNotFound)
			return
		}
	case "IBAN":
		payerAccount, err = db.SelectAccountByIBAN(cfg.DB, abi, req.Payer.AccountIdentification.Value)
		if err != nil {
			fmt.Printf("Error while searching payer account with IBAN %s: %s\n",
				req.Payer.AccountIdentification.Value, err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if payerAccount.Id == "" {
			fmt.Printf("Payer account with IBAN %s not found\n", req.Payer.AccountIdentification.Value)
			w.WriteHeader(http.StatusNotFound)
			return
		}
	}

	if payerAccount.Value.Amount < req.Value.Amount {
		fmt.Printf("Payer account %s without enough funds\n", payerAccount.Id)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Check payee existance on local system
	if req.Payee.AccountIdentification.Type == "IBAN" &&
		req.Payee.AccountIdentification.Value[5:10] == abi {
		_, err = db.SelectAccountByIBAN(cfg.DB, abi, req.Payee.AccountIdentification.Value)
		if err == mongo.ErrNoDocuments {
			fmt.Printf("Payee account %s not found\n", req.Payee.AccountIdentification.Value)
			w.WriteHeader(http.StatusNotFound)
			return
		}
		if err != nil {
			fmt.Printf("Error while searching payee account %s: %s\n",
				req.Payee.AccountIdentification.Value, err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	// Build the new document
	paymentId := primitive.NewObjectID().Hex()
	payment := cha.Payment{
		Id:      paymentId,
		Type:    req.Type,
		Value:   req.Value,
		Payer:   req.Payer,
		Payee:   req.Payee,
		Details: req.Details,
	}

	// Insert the new document
	err = db.InsertPayment(cfg.DB, abi, payment)
	if mongo.IsDuplicateKeyError(err) {
		fmt.Printf("Payment %+v already exists\n", payment)
		w.WriteHeader(http.StatusConflict)
		return
	}
	if err != nil {
		fmt.Printf("Error while inserting payment %+v: %s\n", payment, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Add document to queue
	err = service.QueuePayment(cfg.Queue, abi, payment)
	if err != nil {
		fmt.Printf("Error while queueing payment %+v: %s\n", payment, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Response output
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(payment)
}
