package api

import (
	"bff/config"
	"encoding/json"
	"fmt"
	com "mainframe-lib/common/config"
	dos "mainframe-lib/dossier/model"
	sdos "mainframe-lib/dossier/service"
	"net/http"
	"strconv"
)

func GetOrder(w http.ResponseWriter, r *http.Request) {
	// Extract path parameters
	orderId := r.PathValue(string(config.ContextOrderId))
	if len(orderId) != 24 {
		fmt.Printf("Invalid order id value\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Extract context parameters
	cfg := r.Context().Value(com.ContextConfig).(config.Config)
	auth := r.Context().Value(com.ContextAuth).(string)

	// Get the document
	order, status, err := sdos.GetOrder(cfg.Services.Dossiers, auth, orderId)
	if err != nil {
		fmt.Printf("Error while getting order with id %s: %s\n", orderId, err.Error())
		w.WriteHeader(status)
		return
	}

	// Response output
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Expose-Headers", "*")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(order)
}

func GetOrders(w http.ResponseWriter, r *http.Request) {
	// Extract query parameters
	queryParams := r.URL.Query()
	var err error

	page := 0
	if queryParams.Has(string(config.ContextPage)) {
		page, err = strconv.Atoi(queryParams.Get(string(config.ContextPage)))

		if err != nil {
			fmt.Printf("Invalid %s parameter\n", string(config.ContextPage))
			w.WriteHeader(http.StatusBadRequest)
			return
		}
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

	// Get all documents
	orders, status, err := sdos.GetOrders(cfg.Services.Dossiers, auth, dos.Order{}, page, limit)
	if err != nil {
		fmt.Printf("Error while searching stocks: %s\n", err.Error())
		w.WriteHeader(status)
		return
	}

	// Response output
	if len(orders) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(orders)
}

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	// Parse the request
	var req dos.InsertOrderInput
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
	status, err := sdos.InsertOrder(cfg.Services.Dossiers, auth, req)
	if err != nil {
		fmt.Printf("Error while creating order: %s\n", err.Error())
		w.WriteHeader(status)
		return
	}

	// Response output
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusCreated)
}

func DeleteOrder(w http.ResponseWriter, r *http.Request) {
	// Extract path parameters
	orderId := r.PathValue(string(config.ContextOrderId))
	if len(orderId) != 24 {
		fmt.Printf("Invalid order id value\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Extract context parameters
	cfg := r.Context().Value(com.ContextConfig).(config.Config)
	auth := r.Context().Value(com.ContextAuth).(string)

	// Create a new document
	status, err := sdos.DeleteOrder(cfg.Services.Dossiers, auth, orderId)
	if err != nil {
		fmt.Printf("Error while deleting order: %s\n", err.Error())
		w.WriteHeader(status)
		return
	}

	// Response output
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusNoContent)
}
