package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"mainframe/dossier/config"
	"mainframe/dossier/model"
	"net/http"
)

func GetApiKey(cfg config.Config, apiKeyId string) (model.ApiKey, error) {
	// Construct the request
	url := fmt.Sprintf("http://%s/api-keys/%s", cfg.Services.Security, apiKeyId)

	// Execute the request
	res, err := ExecuteHttpRequest(cfg, http.MethodGet, url, "", "")
	if err != nil {
		return model.ApiKey{}, err
	}

	// Check response
	if res.StatusCode != http.StatusOK {
		return model.ApiKey{}, errors.New("get api key failed")
	}

	// Parse the response
	var apiKey model.ApiKey
	err = json.NewDecoder(res.Body).Decode(&apiKey)
	if err != nil {
		return model.ApiKey{}, errors.New("get api key response convertion failed")
	}

	return apiKey, nil
}
