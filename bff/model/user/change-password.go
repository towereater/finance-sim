package model

type ChangePasswordUserInput struct {
	OldPassword string `json:"password-old"`
	NewPassword string `json:"password-new"`
}
