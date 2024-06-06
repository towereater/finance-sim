package account

import (
	"fmt"
	"net/http"
)

// Get account API function
func AccountOptions(w http.ResponseWriter, r *http.Request) {
	// Response output
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", fmt.Sprintf("%v, %v, %v", http.MethodGet, http.MethodPost, http.MethodOptions))
	w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Credentials, Authorization")
	w.WriteHeader(http.StatusNoContent)
}
