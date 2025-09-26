package service

import (
	"encoding/json"
	"fmt"
	"mainframe-lib/checking-account/model"
	com "mainframe-lib/common/service"
	"net/http"
)

func GetPayment(host string, timeout int, auth string, paymentId string) (model.Payment, int, error) {
	// Construct the request
	url := fmt.Sprintf("http://%s/payments/%s", host, paymentId)

	// Execute the request
	res, err := com.ExecuteHttpRequest(http.MethodGet, url, timeout, auth, "")
	if err != nil {
		return model.Payment{}, http.StatusInternalServerError, err
	}

	// Check response
	if res.StatusCode == http.StatusNotFound {
		return model.Payment{}, res.StatusCode, nil
	}
	if res.StatusCode != http.StatusOK {
		return model.Payment{}, res.StatusCode, fmt.Errorf("get payment returned status %d", res.StatusCode)
	}

	// Parse the response
	var payment model.Payment
	err = json.NewDecoder(res.Body).Decode(&payment)
	if err != nil {
		return model.Payment{}, http.StatusInternalServerError, err
	}

	return payment, res.StatusCode, nil
}

func GetPayments(host string, timeout int, auth string, filter model.Payment, from string, limit int) ([]model.Payment, int, error) {
	// Construct the request
	url := fmt.Sprintf("http://%s/payments", host)

	url = fmt.Sprintf("%s?limit=%d", url, limit)
	if from != "" {
		url = fmt.Sprintf("%s&from=%s", url, from)
	}
	if filter.Payer.AccountIdentification.Type != "" {
		url = fmt.Sprintf("%s&payerType=%s", url, filter.Payer.AccountIdentification.Type)
	}
	if filter.Payer.AccountIdentification.Value != "" {
		url = fmt.Sprintf("%s&payerValue=%s", url, filter.Payer.AccountIdentification.Value)
	}

	// Execute the request
	res, err := com.ExecuteHttpRequest(http.MethodGet, url, timeout, auth, "")
	if err != nil {
		return []model.Payment{}, http.StatusInternalServerError, err
	}

	// Check response
	if res.StatusCode == http.StatusNotFound {
		return []model.Payment{}, res.StatusCode, nil
	}
	if res.StatusCode != http.StatusOK {
		return []model.Payment{}, res.StatusCode, fmt.Errorf("get payments returned status %d", res.StatusCode)
	}

	// Parse the response
	var payments []model.Payment
	err = json.NewDecoder(res.Body).Decode(&payments)
	if err != nil {
		return []model.Payment{}, http.StatusInternalServerError, err
	}

	return payments, res.StatusCode, nil
}

func InsertPayment(host string, timeout int, auth string, payload model.InsertPayment) (model.Payment, int, error) {
	// Construct the request
	url := fmt.Sprintf("http://%s/payments", host)

	// Execute the request
	res, err := com.ExecuteHttpRequest(http.MethodPost, url, timeout, auth, payload)
	if err != nil {
		return model.Payment{}, http.StatusInternalServerError, err
	}

	// Check response
	if res.StatusCode != http.StatusCreated {
		return model.Payment{}, res.StatusCode, fmt.Errorf("insert payment returned status %d", res.StatusCode)
	}

	// Parse the response
	var payment model.Payment
	err = json.NewDecoder(res.Body).Decode(&payment)
	if err != nil {
		return model.Payment{}, http.StatusInternalServerError, err
	}

	return payment, res.StatusCode, nil
}
