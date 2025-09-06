package service

import (
	"encoding/json"
	"fmt"
	"mainframe-lib/checking-account/model"
	com "mainframe-lib/common/service"
	"net/http"
)

func GetPayments(host string, timeout int, auth string, filter model.Payment, from string, limit int) ([]model.Payment, error) {
	// Construct the request
	url := fmt.Sprintf("http://%s/payments", host)

	url = fmt.Sprintf("%s?limit=%d", url, limit)
	if from != "" {
		url = fmt.Sprintf("%s&from=%s", url, from)
	}
	if filter.Payer.AccountId.Account != "" {
		url = fmt.Sprintf("%s&account=%s", url, filter.Payer.AccountId.Account)
	}
	if filter.Payer.AccountId.Service != "" {
		url = fmt.Sprintf("%s&service=%s", url, filter.Payer.AccountId.Service)
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
