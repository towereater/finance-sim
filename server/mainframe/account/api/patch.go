package api

import (
	"bff/api"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"mainframe/account/config"
	"mainframe/account/db"
	"mainframe/account/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Patch account API function
func PatchAccount(w http.ResponseWriter, r *http.Request) {
	// Parsing of the request
	var req model.PatchAccountInput
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Extraction of extra parameters
	id, err := primitive.ObjectIDFromHex(r.PathValue("accountId"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	// Fetch of the account data
	account, err := db.SelectAccount(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Update account in the user accounts list
	if account.Owner != req.Owner {
		err = updateAccountOwner(id.Hex(), account.Owner, req.Owner)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	// Generation of the updated document
	if req.IBAN != "" {
		account.IBAN = req.IBAN
	}
	if req.Owner != "" {
		account.Owner = req.Owner
	}
	if req.Cash != nil {
		account.Cash = *req.Cash
	}

	// Execution of the request
	err = db.UpdateAccount(id, *account)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Response output
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(account)
}

func updateAccountOwner(accountId string, oldOwnerId string, newOwnerId string) error {
	// Construction of the request
	url := "http://" + config.AppConfig.UsersServer.Host + ":" + config.AppConfig.UsersServer.Port
	url = url + "/users/" + oldOwnerId + "/accounts/" + accountId

	// Execution of the request
	res, err := api.ExecuteHttpRequest(http.MethodDelete, url, "")
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// Response parsing
	if res.StatusCode != http.StatusOK {
		body, err := io.ReadAll(res.Body)

		if err != nil {
			return err
		}

		return errors.New(string(body))
	}

	res.Body.Close()

	// Construction of the request
	url = "http://" + config.AppConfig.UsersServer.Host + ":" + config.AppConfig.UsersServer.Port
	url = url + "/users/" + newOwnerId + "/accounts/" + accountId

	// Execution of the request
	res, err = api.ExecuteHttpRequest(http.MethodPost, url, "")
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// Response parsing
	if res.StatusCode != http.StatusOK {
		body, err := io.ReadAll(res.Body)

		if err != nil {
			return err
		}

		return errors.New(string(body))
	}

	return nil
}
