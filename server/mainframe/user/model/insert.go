package model

type InsertUserInput struct {
	Username string   `json:"username"`
	Password string   `json:"password"`
	Name     string   `json:"name"`
	Surname  string   `json:"surname"`
	Birth    string   `json:"birth"`
	Accounts []string `json:"accounts,omitempty"`
}
