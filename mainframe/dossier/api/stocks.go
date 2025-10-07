package api

import (
	"encoding/json"
	"fmt"
	"mainframe/dossier/config"
	"mainframe/dossier/service"
	"net/http"
	"slices"
	"strconv"

	com "mainframe-lib/common/config"
	dos "mainframe-lib/dossier/model"
	ssec "mainframe-lib/security/service"
	xch "mainframe-lib/xchanger/model"
	sxch "mainframe-lib/xchanger/service"
)

func GetStock(w http.ResponseWriter, r *http.Request) {
	// Extract path parameters
	isin := r.PathValue(string(config.ContextIsin))
	if len(isin) != 12 {
		fmt.Printf("Invalid isin value\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Extract context parameters
	cfg := r.Context().Value(com.ContextConfig).(config.Config)
	abi := r.Context().Value(com.ContextAbi).(string)
	auth := r.Context().Value(com.ContextAuth).(string)

	// Get bank data
	bank, status, err := ssec.GetBankByAbi(cfg.Services.Security, auth, abi)
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
	xchangerStock, status, err := sxch.GetStock(cfg.Services.Xchanger, bank.XchangerApiKey, isin)
	if err != nil {
		fmt.Printf("Error while searching isin %s on xchanger: %s\n",
			isin,
			err.Error())
		w.WriteHeader(status)
		return
	}

	// Convert the document to standard format
	stock := service.ToStock(xchangerStock)

	// Response output
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(stock)
}

func GetStocks(w http.ResponseWriter, r *http.Request) {
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
	abi := r.Context().Value(com.ContextAbi).(string)
	auth := r.Context().Value(com.ContextAuth).(string)

	// Get bank data
	bank, status, err := ssec.GetBankByAbi(cfg.Services.Security, auth, abi)
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
	xchangerStocks, status, err := sxch.GetStocks(cfg.Services.Xchanger, bank.XchangerApiKey, xch.Stock{}, page, limit)
	if err != nil {
		fmt.Printf("Error while searching stocks on xchanger: %s\n", err.Error())
		w.WriteHeader(status)
		return
	}
	if len(xchangerStocks) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Convert the document to standard format
	var stocks []dos.Stock
	for _, stock := range slices.All(xchangerStocks) {
		stocks = append(stocks, service.ToStock(stock))
	}

	// Response output
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(stocks)
}
