package model

type LoginUserInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginUserOutput struct {
	Name    string   `json:"name"`
	Surname string   `json:"surname"`
	Birth   string   `json:"birth"`
	Account []string `json:"accounts"`
}
