package api

import (
	"bff/api"
	"errors"
	"io"
	"net/http"

	"mainframe/account/config"
	"mainframe/account/db"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Delete account API function
func DeleteAccount(w http.ResponseWriter, r *http.Request) {
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

	// Delete account from the user accounts list
	err = deleteAccountFromUser(account.Id.Hex(), account.Owner)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Execution of the request
	err = db.DeleteAccount(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Response output
	w.WriteHeader(http.StatusNoContent)
}

func deleteAccountFromUser(accountId string, userId string) error {
	// Construction of the request
	url := "http://" + config.AppConfig.UsersServer.Host + ":" + config.AppConfig.UsersServer.Port
	url = url + "/users/" + userId + "/accounts/" + accountId

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

	return nil
}
