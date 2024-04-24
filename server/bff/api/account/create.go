package account

import (
	"encoding/json"
	"net/http"

	"bff/api"
	cfg "bff/config"
	bff "bff/model/account"

	mf "mainframe/account/model"
)

// Create account API function
func CreateAccount(w http.ResponseWriter, r *http.Request) {
	// Recovery of JWT
	jwt := r.Header.Get("Authorization")
	if jwt == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Construction of the request
	url := "http://" + cfg.AppConfig.Server.Accounts.Host + ":" + cfg.AppConfig.Server.Accounts.Port
	url = url + "/accounts"
	payload := mf.InsertAccountInput{
		Owner: jwt,
	}

	// Execution of the request
	res, err := api.ExecuteHttpRequest(http.MethodPost, url, payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()

	// Response parsing
	var mfAccount mf.Account
	err = json.NewDecoder(res.Body).Decode(&mfAccount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Response generation
	var account = bff.Account{
		Id:    mfAccount.Id,
		IBAN:  mfAccount.IBAN,
		Owner: mfAccount.Owner,
		Cash:  mfAccount.Cash,
	}

	// Response output
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.StatusCode)
	json.NewEncoder(w).Encode(account)
}
