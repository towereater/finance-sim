package model

type User struct {
	Id       string `json:"id" bson:"_id"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	ApiKey   string `json:"apiKey" bson:"apiKey"`
}

type InsertUserInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
	ApiKey   string `json:"apiKey"`
}
