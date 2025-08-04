package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"mainframe/dossier/config"
	"mainframe/dossier/model"
	"net/http"
)

func InsertXChangerDossier(cfg config.Config, payload model.InsertXChangerDossierInput) (model.XChangerDossier, error) {
	// Construct the request
	url := fmt.Sprintf("http://%s/dossiers", cfg.Services.Xchanger.Host)

	// Execute the request
	res, err := ExecuteHttpRequest(cfg, http.MethodPost, url, payload, cfg.Services.Xchanger.ApiKey)
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

func DeleteXChangerDossier(cfg config.Config, dossierId string) error {
	// Construct the request
	url := fmt.Sprintf("http://%s/dossiers/%s",
		cfg.Services.Xchanger.Host, dossierId)

	// Execute the request
	res, err := ExecuteHttpRequest(cfg, http.MethodDelete, url, "", cfg.Services.Xchanger.ApiKey)
	if err != nil {
		return err
	}

	// Check response
	if res.StatusCode != http.StatusNoContent {
		return errors.New("delete xchanger dossier failed")
	}

	return nil
}
