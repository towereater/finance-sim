package service

import (
	"encoding/json"
	"fmt"
	"mainframe-lib/checking-account/model"
	ccom "mainframe-lib/common/config"
	scom "mainframe-lib/common/service"
	"net/http"
)

func GetPayment(service ccom.Service, auth string, paymentId string) (model.Payment, int, error) {
	// Construct the request
	url := fmt.Sprintf("/payments/%s", paymentId)

	// Execute the request
	res, err := scom.ExecuteHttpRequest(service, http.MethodGet, url, auth, "")
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

func GetPayments(service ccom.Service, auth string, filter model.Payment, from string, limit int) ([]model.Payment, int, error) {
	// Construct the request
	url := "/payments"

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
	res, err := scom.ExecuteHttpRequest(service, http.MethodGet, url, auth, "")
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

func InsertPayment(service ccom.Service, auth string, payload model.InsertPayment) (model.Payment, int, error) {
	// Construct the request
	url := "/payments"

	// Execute the request
	res, err := scom.ExecuteHttpRequest(service, http.MethodPost, url, auth, payload)
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
