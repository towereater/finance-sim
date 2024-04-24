package api

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"mainframe/account/config"
	"mainframe/account/db"
	"mainframe/account/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Insert account API function
func InsertAccount(w http.ResponseWriter, r *http.Request) {
	// Parsing of the request
	var req model.InsertAccountInput
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Generation of the new document
	account := model.Account{
		Id:    primitive.NewObjectID(),
		IBAN:  primitive.NewObjectID().Hex(),
		Owner: req.Owner,
		Cash:  0,
	}

	// Execution of the request
	err = db.InsertAccount(account)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Insert account in the user accounts list
	err = insertAccountToUser(account.Id.Hex(), req.Owner)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Response output
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(account)
}

func insertAccountToUser(accountId string, userId string) error {
	// Construction of the request
	url := "http://" + config.AppConfig.UsersServer.Host + ":" + config.AppConfig.UsersServer.Port
	url = url + "/users/" + userId + "/accounts/" + accountId

	// Execution of the request
	res, err := ExecuteHttpRequest(http.MethodPost, url, "")
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
