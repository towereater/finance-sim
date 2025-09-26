package service

import (
	"encoding/json"
	"fmt"
	"mainframe-lib/dossier/model"

	com "mainframe-lib/common/service"
	"net/http"
)

func GetOrder(host string, timeout int, auth string, orderId string) (model.Order, int, error) {
	// Construct the request
	url := fmt.Sprintf("http://%s/orders/%s", host, orderId)

	// Execute the request
	res, err := com.ExecuteHttpRequest(http.MethodGet, url, timeout, auth, "")
	if err != nil {
		return model.Order{}, http.StatusInternalServerError, err
	}

	// Check response
	if res.StatusCode == http.StatusNotFound {
		return model.Order{}, res.StatusCode, nil
	}
	if res.StatusCode != http.StatusOK {
		return model.Order{}, res.StatusCode, fmt.Errorf("get order returned status %d", res.StatusCode)
	}

	// Parse the response
	var order model.Order
	err = json.NewDecoder(res.Body).Decode(&order)
	if err != nil {
		return model.Order{}, http.StatusInternalServerError, err
	}

	return order, res.StatusCode, nil
}

func GetOrders(host string, timeout int, auth string, filter model.Order, page int, limit int) ([]model.Order, int, error) {
	// Construct the request
	url := fmt.Sprintf("http://%s/orders", host)

	url = fmt.Sprintf("%s?limit=%d", url, limit)
	url = fmt.Sprintf("%s&page=%d", url, page)

	// Execute the request
	res, err := com.ExecuteHttpRequest(http.MethodGet, url, timeout, auth, "")
	if err != nil {
		return []model.Order{}, http.StatusInternalServerError, err
	}

	// Check response
	if res.StatusCode == http.StatusNotFound {
		return []model.Order{}, res.StatusCode, nil
	}
	if res.StatusCode != http.StatusOK {
		return []model.Order{}, res.StatusCode, fmt.Errorf("get orders returned status %d", res.StatusCode)
	}

	// Parse the response
	var orders []model.Order
	err = json.NewDecoder(res.Body).Decode(&orders)
	if err != nil {
		return []model.Order{}, http.StatusInternalServerError, err
	}

	return orders, res.StatusCode, nil
}

func InsertOrder(host string, timeout int, auth string, payload model.InsertOrderInput) (int, error) {
	// Construct the request
	url := fmt.Sprintf("http://%s/orders", host)

	// Execute the request
	res, err := com.ExecuteHttpRequest(http.MethodPost, url, timeout, auth, payload)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	// Check response
	if res.StatusCode != http.StatusCreated {
		return res.StatusCode, fmt.Errorf("insert order returned status %d", res.StatusCode)
	}

	return res.StatusCode, nil
}

func DeleteOrder(host string, timeout int, auth string, orderId string) (int, error) {
	// Construct the request
	url := fmt.Sprintf("http://%s/orders/%s", host, orderId)

	// Execute the request
	res, err := com.ExecuteHttpRequest(http.MethodDelete, url, timeout, auth, "")
	if err != nil {
		return http.StatusInternalServerError, err
	}

	// Check response
	if res.StatusCode != http.StatusNoContent {
		return res.StatusCode, fmt.Errorf("delete order returned status %d", res.StatusCode)
	}

	return res.StatusCode, nil
}
