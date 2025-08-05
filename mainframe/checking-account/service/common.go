package service

import (
	"bytes"
	"mainframe/checking-account/config"
	"time"

	"encoding/json"
	"net/http"
)

func ExecuteHttpRequest(cfg config.Config, method string, url string, apiKey string, payload any) (*http.Response, error) {
	// Convert the payload
	bytesPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	// Construct the request
	req, err := http.NewRequest(method, url, bytes.NewBuffer(bytesPayload))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	if apiKey != "" {
		req.Header.Set("Authorization", apiKey)
	}

	client := &http.Client{
		Timeout: time.Duration(cfg.Services.Timeout) * time.Second,
	}

	// Execute the request
	return client.Do(req)
}
