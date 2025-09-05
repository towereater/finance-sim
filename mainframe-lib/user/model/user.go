package model

type User struct {
	Id       string    `json:"id" bson:"_id"`
	Username string    `json:"username" bson:"username"`
	Password string    `json:"password" bson:"password"`
	Name     string    `json:"name" bson:"name"`
	Surname  string    `json:"surname" bson:"surname"`
	Birth    string    `json:"birth" bson:"birth"`
	Cab      string    `json:"cab"`
	Accounts []Account `json:"accounts,omitempty" bson:"accounts,omitempty"`
}

type InsertUserInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Birth    string `json:"birth"`
	Cab      string `json:"cab"`
}

type UpdateUserInput struct {
	Password string `json:"password,omitempty"`
}
