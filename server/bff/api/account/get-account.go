package account

import (
	"encoding/json"
	"net/http"

	"bff/api"
	cfg "bff/config"
	bff "bff/model/account"

	mf "mainframe/account/model"
)

// Get account API function
func GetAccount(w http.ResponseWriter, r *http.Request) {
	// Extraction of extra parameters
	accountId := r.PathValue("accountId")

	// Recovery of JWT
	jwt := r.Header.Get("Authorization")
	if jwt == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Construction of the request
	url := "http://" + cfg.AppConfig.Server.Accounts.Host + ":" + cfg.AppConfig.Server.Accounts.Port
	url = url + "/accounts/" + accountId

	// Execution of the request
	res, err := api.ExecuteHttpRequest(http.MethodGet, url, "")
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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.StatusCode)
	json.NewEncoder(w).Encode(account)
}
