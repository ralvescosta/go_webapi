package dtos

type AuthenticationDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
