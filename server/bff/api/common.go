package api

import (
	"bytes"
	"time"

	"encoding/json"
	"net/http"
)

func ExecuteHttpRequest(method string, url string, payload any) (*http.Response, error) {
	// Convertion of the payload
	bytesPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	// Construction of the request
	req, err := http.NewRequest(method, url, bytes.NewBuffer(bytesPayload))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	// Execution of the request
	return client.Do(req)
}
