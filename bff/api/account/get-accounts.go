package account

import (
	"encoding/json"
	"net/http"

	"bff/api"
	cfg "bff/config"
	bff "bff/model/account"

	mf "mainframe/account/model"
)

// Get accounts API function
func GetAccounts(w http.ResponseWriter, r *http.Request) {
	// Recovery of JWT
	jwt := r.Header.Get("Authorization")
	if jwt == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Construction of the request
	url := "http://" + cfg.AppConfig.Server.Accounts.Host + ":" + cfg.AppConfig.Server.Accounts.Port
	url = url + "/accounts?owner=" + jwt

	// Execution of the request
	res, err := api.ExecuteHttpRequest(http.MethodGet, url, "")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if res.StatusCode == http.StatusNoContent {
		w.WriteHeader(res.StatusCode)
		return
	}

	defer res.Body.Close()

	// Response parsing
	var mfAccounts []mf.Account
	err = json.NewDecoder(res.Body).Decode(&mfAccounts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Response generation
	var accounts []bff.AccountList
	for _, a := range mfAccounts {
		accounts = append(accounts, bff.AccountList{
			Id:    a.Id,
			IBAN:  a.IBAN,
			Owner: a.Owner,
		})
	}

	response := struct {
		Accounts []bff.AccountList `json:"accounts"`
	}{
		Accounts: accounts,
	}

	// Response output
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.StatusCode)
	json.NewEncoder(w).Encode(response)
}
