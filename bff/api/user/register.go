package user

import (
	"encoding/json"
	"net/http"

	"bff/api"
	"bff/config"
	bff "bff/model/user"

	mf "mainframe/user/model"
)

// Register user API function
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	// Parsing of the request
	var req bff.RegisterUserInput
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Construction of the request
	url := "http://" + config.AppConfig.Server.Users.Host + ":" + config.AppConfig.Server.Users.Port + "/users"
	payload := mf.InsertUserInput{
		Username: req.Username,
		Password: req.Password,
		Name:     req.Name,
		Surname:  req.Surname,
		Birth:    req.Birth,
	}

	// Execution of the request
	res, err := api.ExecuteHttpRequest(http.MethodPost, url, payload)
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

	// Response output
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.StatusCode)
	json.NewEncoder(w).Encode(user)
}
