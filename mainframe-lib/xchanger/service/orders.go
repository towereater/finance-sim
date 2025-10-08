package service

import (
	"encoding/json"
	"fmt"
	ccom "mainframe-lib/common/config"
	scom "mainframe-lib/common/service"
	"mainframe-lib/xchanger/model"
	"net/http"
)

func GetOrder(service ccom.Service, auth string, orderId string) (model.Order, int, error) {
	// Construct the request
	url := fmt.Sprintf("/orders/%s", orderId)

	// Execute the request
	res, err := scom.ExecuteHttpRequest(service, http.MethodGet, url, auth, "")
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

func GetOrders(service ccom.Service, auth string, filter model.Order, page int, size int) ([]model.Order, int, error) {
	// Construct the request
	url := "/orders"

	url = fmt.Sprintf("%s?page=%d", url, page)
	url = fmt.Sprintf("%s&size=%d", url, size)

	// Execute the request
	res, err := scom.ExecuteHttpRequest(service, http.MethodGet, url, auth, "")
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

func InsertOrder(service ccom.Service, auth string, payload model.InsertOrderInput) (model.Order, int, error) {
	// Construct the request
	url := "/orders"

	// Execute the request
	res, err := scom.ExecuteHttpRequest(service, http.MethodPost, url, auth, payload)
	if err != nil {
		return model.Order{}, http.StatusInternalServerError, err
	}

	// Check response
	if res.StatusCode != http.StatusCreated {
		return model.Order{}, res.StatusCode, fmt.Errorf("insert xchanger order failed")
	}

	// Parse the response
	var order model.Order
	err = json.NewDecoder(res.Body).Decode(&order)
	if err != nil {
		return model.Order{}, http.StatusInternalServerError, fmt.Errorf("insert xchanger order response convertion failed")
	}

	return order, res.StatusCode, nil
}

func DeleteOrder(service ccom.Service, auth string, orderId string) (int, error) {
	// Construct the request
	url := fmt.Sprintf("/orders/%s", orderId)

	// Execute the request
	res, err := scom.ExecuteHttpRequest(service, http.MethodDelete, url, auth, "")
	if err != nil {
		return http.StatusInternalServerError, err
	}

	// Check response
	if res.StatusCode != http.StatusNoContent {
		return res.StatusCode, fmt.Errorf("delete xchanger order failed")
	}

	return res.StatusCode, nil
}
