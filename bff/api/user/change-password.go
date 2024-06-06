package user

import (
	"encoding/json"
	"net/http"

	"bff/api"
	"bff/config"
	bff "bff/model/user"

	mf "mainframe/user/model"
)

// Change password user API function
func ChangePasswordUser(w http.ResponseWriter, r *http.Request) {
	// Parsing of the request
	var changePasswordUserInput bff.ChangePasswordUserInput
	err := json.NewDecoder(r.Body).Decode(&changePasswordUserInput)
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
	if changePasswordUserInput.OldPassword != user.Password {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Construction of the request
	payload := mf.PatchUserInput{
		Password: changePasswordUserInput.NewPassword,
	}

	// Execution of the request
	res, err = api.ExecuteHttpRequest(http.MethodPatch, url, payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()

	// Response output
	w.WriteHeader(res.StatusCode)
}
