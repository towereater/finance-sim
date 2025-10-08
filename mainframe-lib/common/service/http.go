package service

import (
	"bytes"
	"encoding/json"
	"mainframe-lib/common/config"
	"net/http"
	"time"
)

func ExecuteHttpRequest(service config.Service, method string, url string, auth string, payload any) (*http.Response, error) {
	// Convert the payload
	bytesPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	// Construct the request
	req, err := http.NewRequest(method, service.Host+url, bytes.NewBuffer(bytesPayload))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", auth)

	client := &http.Client{
		Timeout: time.Duration(service.Timeout) * time.Second,
	}

	// Execute the request
	return client.Do(req)
}
