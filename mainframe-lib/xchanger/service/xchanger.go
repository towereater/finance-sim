package service

import (
	"encoding/json"
	"errors"
	"fmt"
	com "mainframe-lib/common/service"
	"mainframe-lib/xchanger/model"
	"net/http"
)

func InsertXChangerDossier(host string, timeout int, auth string, payload model.InsertXChangerDossierInput) (model.XChangerDossier, error) {
	// Construct the request
	url := fmt.Sprintf("http://%s/dossiers", host)

	// Execute the request
	res, err := com.ExecuteHttpRequest(http.MethodPost, url, timeout, auth, payload)
	if err != nil {
		return model.XChangerDossier{}, err
	}

	// Check response
	if res.StatusCode != http.StatusCreated {
		return model.XChangerDossier{}, errors.New("insert xchanger dossier failed")
	}

	// Parse the response
	var xchangerDossier model.XChangerDossier
	err = json.NewDecoder(res.Body).Decode(&xchangerDossier)
	if err != nil {
		return model.XChangerDossier{}, errors.New("insert xchanger dossier response convertion failed")
	}

	return xchangerDossier, nil
}

func DeleteXChangerDossier(host string, timeout int, auth string, dossierId string) error {
	// Construct the request
	url := fmt.Sprintf("http://%s/dossiers/%s", host, dossierId)

	// Execute the request
	res, err := com.ExecuteHttpRequest(http.MethodDelete, url, timeout, auth, "")
	if err != nil {
		return err
	}

	// Check response
	if res.StatusCode != http.StatusNoContent {
		return errors.New("delete xchanger dossier failed")
	}

	return nil
}
