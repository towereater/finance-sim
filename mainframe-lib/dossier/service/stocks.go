package service

import (
	"encoding/json"
	"fmt"
	com "mainframe-lib/common/service"
	"mainframe-lib/dossier/model"
	"net/http"
)

func GetStock(host string, timeout int, auth string, isin string) (model.Stock, int, error) {
	// Construct the request
	url := fmt.Sprintf("http://%s/stocks/%s", host, isin)

	// Execute the request
	res, err := com.ExecuteHttpRequest(http.MethodGet, url, timeout, auth, "")
	if err != nil {
		return model.Stock{}, http.StatusInternalServerError, err
	}

	// Check response
	if res.StatusCode == http.StatusNotFound {
		return model.Stock{}, res.StatusCode, nil
	}
	if res.StatusCode != http.StatusOK {
		return model.Stock{}, res.StatusCode, fmt.Errorf("get stock returned status %d", res.StatusCode)
	}

	// Parse the response
	var stock model.Stock
	err = json.NewDecoder(res.Body).Decode(&stock)
	if err != nil {
		return model.Stock{}, http.StatusInternalServerError, err
	}

	return stock, res.StatusCode, nil
}

func GetStocks(host string, timeout int, auth string, filter model.Stock, page int, limit int) ([]model.Stock, int, error) {
	// Construct the request
	url := fmt.Sprintf("http://%s/stocks", host)

	url = fmt.Sprintf("%s?limit=%d", url, limit)
	url = fmt.Sprintf("%s&page=%d", url, page)

	// Execute the request
	res, err := com.ExecuteHttpRequest(http.MethodGet, url, timeout, auth, "")
	if err != nil {
		return []model.Stock{}, http.StatusInternalServerError, err
	}

	// Check response
	if res.StatusCode == http.StatusNotFound {
		return []model.Stock{}, res.StatusCode, nil
	}
	if res.StatusCode != http.StatusOK {
		return []model.Stock{}, res.StatusCode, fmt.Errorf("get stocks returned status %d", res.StatusCode)
	}

	// Parse the response
	var stocks []model.Stock
	err = json.NewDecoder(res.Body).Decode(&stocks)
	if err != nil {
		return []model.Stock{}, http.StatusInternalServerError, err
	}

	return stocks, res.StatusCode, nil
}
