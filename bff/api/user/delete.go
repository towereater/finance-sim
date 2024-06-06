package user

import (
	"encoding/json"
	mf "mainframe/user/model"
	"net/http"

	"bff/api"
	"bff/config"
	bff "bff/model/user"
)

// Delete user API function
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Parsing of the request
	var deleteUserInput bff.DeleteUserInput
	err := json.NewDecoder(r.Body).Decode(&deleteUserInput)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Recovery of JWT
	jwt := r.Header.Get("Authorization")
	if jwt == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Construction of the request
	url := "http://" + config.AppConfig.Server.Users.Host + ":" + config.AppConfig.Server.Users.Port + "/users/" + jwt

	// Execution of the request
	res, err := api.ExecuteHttpRequest(http.MethodGet, url, "")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()

	// Response parsing
	var user mf.User
	err = json.NewDecoder(res.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Password confirmation
	if deleteUserInput.Password != user.Password {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Execution of the request
	res, err = api.ExecuteHttpRequest(http.MethodDelete, url, "")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()

	// Response output
	w.WriteHeader(res.StatusCode)
}
