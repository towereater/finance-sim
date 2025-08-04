package model

type AddAccountToUserInput struct {
	Id AccountId `json:"id"`
}

type RemoveAccountFromUserInput struct {
	Id AccountId `json:"id"`
}
