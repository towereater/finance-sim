package api

import (
	"encoding/json"
	"fmt"
	"mainframe/dossier/config"
	"mainframe/dossier/db"
	"mainframe/dossier/service"
	"net/http"
	"slices"
	"strconv"

	com "mainframe-lib/common/config"
	dos "mainframe-lib/dossier/model"
	ssec "mainframe-lib/security/service"
	xch "mainframe-lib/xchanger/model"
	sxch "mainframe-lib/xchanger/service"

	"go.mongodb.org/mongo-driver/mongo"
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
	abi := r.Context().Value(com.ContextAbi).(string)
	auth := r.Context().Value(com.ContextAuth).(string)

	// Get bank data
	bank, status, err := ssec.GetBankByAbi(cfg.Services.Security, cfg.Services.Timeout, auth, abi)
	if err != nil {
		fmt.Printf("Error while searching bank with abi %s: %s\n", abi, err.Error())
		w.WriteHeader(status)
		return
	}
	if bank.XchangerApiKey == "" {
		fmt.Printf("Bank with abi %s does not have access to xchanger\n", abi)
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// Select the document
	xchangerOrder, status, err := sxch.GetOrder(cfg.Services.Xchanger, cfg.Services.Timeout, bank.XchangerApiKey, orderId)
	if err != nil {
		fmt.Printf("Error while searching order %s on xchanger: %s\n",
			orderId,
			err.Error())
		w.WriteHeader(status)
		return
	}

	// Convert the document to standard format
	order := service.ToOrder(xchangerOrder)

	// Response output
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

		if page > 50 {
			page = 50
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
	abi := r.Context().Value(com.ContextAbi).(string)
	auth := r.Context().Value(com.ContextAuth).(string)

	// Get bank data
	bank, status, err := ssec.GetBankByAbi(cfg.Services.Security, cfg.Services.Timeout, auth, abi)
	if err != nil {
		fmt.Printf("Error while searching bank with abi %s: %s\n", abi, err.Error())
		w.WriteHeader(status)
		return
	}
	if bank.XchangerApiKey == "" {
		fmt.Printf("Bank with abi %s does not have access to xchanger\n", abi)
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// Select all documents
	xchangerOrders, status, err := sxch.GetOrders(cfg.Services.Xchanger, cfg.Services.Timeout, bank.XchangerApiKey, xch.Order{}, page, limit)
	if err != nil {
		fmt.Printf("Error while searching orders on xchanger: %s\n", err.Error())
		w.WriteHeader(status)
		return
	}
	if len(xchangerOrders) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Convert the document to standard format
	var orders []dos.Order
	for _, order := range slices.All(xchangerOrders) {
		orders = append(orders, service.ToOrder(order))
	}

	// Response output
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(orders)
}

func InsertOrder(w http.ResponseWriter, r *http.Request) {
	// Parse the request
	var req dos.InsertOrderInput
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		fmt.Printf("Could not convert request body\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if len(req.Dossier) != 24 {
		fmt.Printf("Invalid dossier owner\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if len(req.Isin) != 12 {
		fmt.Printf("Invalid dossier checking account\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Extract context parameters
	cfg := r.Context().Value(com.ContextConfig).(config.Config)
	abi := r.Context().Value(com.ContextAbi).(string)
	auth := r.Context().Value(com.ContextAuth).(string)

	// Get bank data
	bank, status, err := ssec.GetBankByAbi(cfg.Services.Security, cfg.Services.Timeout, auth, abi)
	if err != nil {
		fmt.Printf("Error while searching bank with abi %s: %s\n", abi, err.Error())
		w.WriteHeader(status)
		return
	}
	if bank.XchangerApiKey == "" {
		fmt.Printf("Bank with abi %s does not have access to xchanger\n", abi)
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// Get dossier data
	dossier, err := db.SelectDossier(cfg, abi, req.Dossier)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No dossiers with id %s\n", req.Dossier)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		fmt.Printf("Error while searching dossier with id %s: %s\n", req.Dossier, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Create a new order on xchanger
	xchangerPayload := xch.InsertOrderInput{
		Dossier: dossier.XChangerDossier,
		Isin:    req.Isin,
		Type:    req.Type,
		Price: xch.Price{
			Amount:   req.Price.Amount,
			Currency: req.Price.Currency,
		},
		Quantity: req.Quantity,
		Options:  req.Options,
	}

	xchangerOrder, status, err := sxch.InsertOrder(cfg.Services.Xchanger, cfg.Services.Timeout, bank.XchangerApiKey, xchangerPayload)
	if err != nil {
		fmt.Printf("Error while creating order %+v on xchanger: %s\n",
			xchangerPayload,
			err.Error())
		w.WriteHeader(status)
		return
	}

	// Convert the document to standard format
	order := service.ToOrder(xchangerOrder)

	// Response output
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)
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
	abi := r.Context().Value(com.ContextAbi).(string)
	auth := r.Context().Value(com.ContextAuth).(string)

	// Get bank data
	bank, status, err := ssec.GetBankByAbi(cfg.Services.Security, cfg.Services.Timeout, auth, abi)
	if err != nil {
		fmt.Printf("Error while searching bank with abi %s: %s\n", abi, err.Error())
		w.WriteHeader(status)
		return
	}
	if bank.XchangerApiKey == "" {
		fmt.Printf("Bank with abi %s does not have access to xchanger\n", abi)
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// Delete order from xchanger
	status, err = sxch.DeleteOrder(cfg.Services.Xchanger, cfg.Services.Timeout, bank.XchangerApiKey, orderId)
	if err != nil {
		fmt.Printf("Error while deleting xchanger order %s: %s\n",
			orderId,
			err.Error())

		w.WriteHeader(status)
		return
	}

	// Response output
	w.WriteHeader(http.StatusNoContent)
}
