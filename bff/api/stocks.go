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
	auth := r.Context().Value(com.ContextAuth).(string)

	// Get the document
	stock, status, err := sdos.GetStock(cfg.Services.Dossiers, auth, isin)
	if err != nil {
		fmt.Printf("Error while getting stock with isin %s: %s\n", isin, err.Error())
		w.WriteHeader(status)
		return
	}

	// Response output
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Expose-Headers", "*")
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
	auth := r.Context().Value(com.ContextAuth).(string)

	// Get all documents
	stocks, status, err := sdos.GetStocks(cfg.Services.Dossiers, auth, dos.Stock{}, page, limit)
	if err != nil {
		fmt.Printf("Error while searching stocks: %s\n", err.Error())
		w.WriteHeader(status)
		return
	}

	// Response output
	if len(stocks) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(stocks)
}
