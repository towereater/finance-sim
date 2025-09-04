package model

type GetUserInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type GetUserOutput struct {
	Name          string `json:"name"`
	Surname       string `json:"surname"`
	Authorization string `json:"authorization"`
}
