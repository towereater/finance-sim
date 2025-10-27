package service

import (
	"encoding/json"
	"fmt"
	ccom "mainframe-lib/common/config"
	scom "mainframe-lib/common/service"
	"mainframe-lib/xchanger/model"
	"net/http"
)

func GetDossier(service ccom.Service, auth string, dossierId string) (model.Dossier, int, error) {
	// Construct the request
	url := fmt.Sprintf("/dossiers/%s", dossierId)

	// Execute the request
	res, err := scom.ExecuteHttpRequest(service, http.MethodGet, url, auth, "")
	if err != nil {
		return model.Dossier{}, http.StatusInternalServerError, err
	}

	// Check response
	if res.StatusCode == http.StatusNotFound {
		return model.Dossier{}, res.StatusCode, nil
	}
	if res.StatusCode != http.StatusOK {
		return model.Dossier{}, res.StatusCode, fmt.Errorf("get dossier returned status %d", res.StatusCode)
	}

	// Parse the response
	var dossier model.Dossier
	err = json.NewDecoder(res.Body).Decode(&dossier)
	if err != nil {
		return model.Dossier{}, http.StatusInternalServerError, err
	}

	return dossier, res.StatusCode, nil
}

func InsertDossier(service ccom.Service, auth string, payload model.InsertDossierInput) (model.Dossier, int, error) {
	// Construct the request
	url := "/dossiers"

	// Execute the request
	res, err := scom.ExecuteHttpRequest(service, http.MethodPost, url, auth, payload)
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

func DeleteDossier(service ccom.Service, auth string, dossierId string) (int, error) {
	// Construct the request
	url := fmt.Sprintf("/dossiers/%s", dossierId)

	// Execute the request
	res, err := scom.ExecuteHttpRequest(service, http.MethodDelete, url, auth, "")
	if err != nil {
		return http.StatusInternalServerError, err
	}

	// Check response
	if res.StatusCode != http.StatusNoContent {
		return res.StatusCode, fmt.Errorf("delete xchanger dossier failed")
	}

	return res.StatusCode, nil
}
