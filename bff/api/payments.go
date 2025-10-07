package api

import (
	"bff/config"
	"encoding/json"
	"fmt"
	cha "mainframe-lib/checking-account/model"
	scha "mainframe-lib/checking-account/service"
	com "mainframe-lib/common/config"
	"net/http"
	"strconv"
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
	auth := r.Context().Value(com.ContextAuth).(string)

	// Get the document
	payment, status, err := scha.GetPayment(cfg.Services.CheckingAccounts, auth, paymentId)
	if err != nil {
		fmt.Printf("Error while getting payment: %s\n", err.Error())
		w.WriteHeader(status)
		return
	}

	// Response output
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Expose-Headers", "*")
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

	// Extract context parameters
	cfg := r.Context().Value(com.ContextConfig).(config.Config)
	auth := r.Context().Value(com.ContextAuth).(string)

	// Build the filter
	var filter cha.Payment
	filter.Type = queryParams.Get("paymentType")
	filter.Payer.AccountIdentification.Type = queryParams.Get("payerType")
	filter.Payer.AccountIdentification.Value = queryParams.Get("payerValue")

	// Get all documents
	payments, status, err := scha.GetPayments(cfg.Services.CheckingAccounts, auth, filter, from, limit)
	if err != nil {
		fmt.Printf("Error while searching payments with filter %+v: %s\n",
			filter, err.Error())
		w.WriteHeader(status)
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

func CreatePayment(w http.ResponseWriter, r *http.Request) {
	// Parse the request
	var req cha.InsertPayment
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		fmt.Printf("Could not convert request body\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Extract context parameters
	cfg := r.Context().Value(com.ContextConfig).(config.Config)
	auth := r.Context().Value(com.ContextAuth).(string)

	// Create a new document
	payment, status, err := scha.InsertPayment(cfg.Services.CheckingAccounts, auth, req)
	if err != nil {
		fmt.Printf("Error while creating payment: %s\n", err.Error())
		w.WriteHeader(status)
		return
	}

	// Response output
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(payment)
}
