package account

import (
	"net/http"

	"bff/api"
	cfg "bff/config"
)

// Delete account API function
func DeleteAccount(w http.ResponseWriter, r *http.Request) {
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
	res, err := api.ExecuteHttpRequest(http.MethodDelete, url, "")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Response output
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.StatusCode)
}
