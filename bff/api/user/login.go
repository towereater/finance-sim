package user

import (
	"encoding/json"
	"net/http"

	"bff/api"
	"bff/config"
	bff "bff/model/user"

	mf "mainframe/user/model"
)

// Login user API function
func LoginUser(w http.ResponseWriter, r *http.Request) {
	// Parsing of the request
	var req bff.LoginUserInput
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Construction of the request
	url := "http://" + config.AppConfig.Server.Users.Host + ":" + config.AppConfig.Server.Users.Port + "/users"
	queryParameters := "?username=" + req.Username + "&password=" + req.Password

	// Execution of the request
	res, err := api.ExecuteHttpRequest(http.MethodGet, url+queryParameters, "")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()

	// Response validation
	if res.StatusCode != http.StatusOK {
		http.Error(w, "Invalid credentials", http.StatusNotFound)
		return
	}

	// Response parsing
	var user []mf.User
	err = json.NewDecoder(res.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(user) > 1 {
		http.Error(w, "More than one user with the same username found", http.StatusInternalServerError)
		return
	}

	loginUserOutput := bff.LoginUserOutput{
		Name:    user[0].Name,
		Surname: user[0].Surname,
		Birth:   user[0].Birth,
		Account: user[0].Accounts,
	}

	// Response output
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Expose-Headers", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Jwt", user[0].Id.Hex())
	w.WriteHeader(res.StatusCode)
	json.NewEncoder(w).Encode(loginUserOutput)
}
