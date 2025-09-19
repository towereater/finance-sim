package service

import (
	"encoding/json"
	"fmt"
	"mainframe-lib/checking-account/model"
	com "mainframe-lib/common/service"
	"net/http"
)

func GetPayment(host string, timeout int, auth string, paymentId string) (model.Payment, error) {
	// Construct the request
	url := fmt.Sprintf("http://%s/payments/%s", host, paymentId)

	// Execute the request
	res, err := com.ExecuteHttpRequest(http.MethodGet, url, timeout, auth, "")
	if err != nil {
		return model.Payment{}, err
	}

	// Check response
	if res.StatusCode != http.StatusOK {
		return model.Payment{}, fmt.Errorf("get payment returned status %d", res.StatusCode)
	}

	// Parse the response
	var payment model.Payment
	err = json.NewDecoder(res.Body).Decode(&payment)
	if err != nil {
		return model.Payment{}, err
	}

	return payment, nil
}

func GetPayments(host string, timeout int, auth string, filter model.Payment, from string, limit int) ([]model.Payment, error) {
	// Construct the request
	url := fmt.Sprintf("http://%s/payments", host)

	url = fmt.Sprintf("%s?limit=%d", url, limit)
	if from != "" {
		url = fmt.Sprintf("%s&from=%s", url, from)
	}
	if filter.Payer.Account != "" {
		url = fmt.Sprintf("%s&account=%s", url, filter.Payer.Account)
	}

	// Execute the request
	res, err := com.ExecuteHttpRequest(http.MethodGet, url, timeout, auth, "")
	if err != nil {
		return []model.Payment{}, err
	}

	// Check response
	if res.StatusCode != http.StatusOK {
		return []model.Payment{}, fmt.Errorf("get payments returned status %d", res.StatusCode)
	}

	// Parse the response
	var payments []model.Payment
	err = json.NewDecoder(res.Body).Decode(&payments)
	if err != nil {
		return []model.Payment{}, err
	}

	return payments, nil
}

func InsertPayment(host string, timeout int, auth string, payload model.InsertPayment) (model.Payment, error) {
	// Construct the request
	url := fmt.Sprintf("http://%s/payments", host)

	// Execute the request
	res, err := com.ExecuteHttpRequest(http.MethodPost, url, timeout, auth, payload)
	if err != nil {
		return model.Payment{}, err
	}

	// Check response
	if res.StatusCode != http.StatusCreated {
		return model.Payment{}, fmt.Errorf("insert payment returned status %d", res.StatusCode)
	}

	// Parse the response
	var payment model.Payment
	err = json.NewDecoder(res.Body).Decode(&payment)
	if err != nil {
		return model.Payment{}, err
	}

	return payment, nil
}
