package service

import (
	"encoding/json"
	"fmt"
	com "mainframe-lib/common/service"
	"mainframe-lib/xchanger/model"
	"net/http"
)

func InsertDossier(host string, timeout int, auth string, payload model.InsertDossierInput) (model.Dossier, int, error) {
	// Construct the request
	url := fmt.Sprintf("http://%s/dossiers", host)

	// Execute the request
	res, err := com.ExecuteHttpRequest(http.MethodPost, url, timeout, auth, payload)
	if err != nil {
		return model.Dossier{}, http.StatusInternalServerError, err
	}

	// Check response
	if res.StatusCode != http.StatusCreated {
		return model.Dossier{}, res.StatusCode, fmt.Errorf("insert xchanger dossier failed")
	}

	// Parse the response
	var dossier model.Dossier
	err = json.NewDecoder(res.Body).Decode(&dossier)
	if err != nil {
		return model.Dossier{}, http.StatusInternalServerError, fmt.Errorf("insert xchanger dossier response convertion failed")
	}

	return dossier, res.StatusCode, nil
}

func DeleteDossier(host string, timeout int, auth string, dossierId string) (int, error) {
	// Construct the request
	url := fmt.Sprintf("http://%s/dossiers/%s", host, dossierId)

	// Execute the request
	res, err := com.ExecuteHttpRequest(http.MethodDelete, url, timeout, auth, "")
	if err != nil {
		return http.StatusInternalServerError, err
	}

	// Check response
	if res.StatusCode != http.StatusNoContent {
		return res.StatusCode, fmt.Errorf("delete xchanger dossier failed")
	}

	return res.StatusCode, nil
}
