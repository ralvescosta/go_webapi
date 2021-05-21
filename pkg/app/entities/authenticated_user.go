package entities

import "time"

type AuthenticatedUser struct {
	Id          uint32    `json:"id"`
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	Email       string    `json:"email"`
	AccessToken string    `json:"accessToken"`
	ExpireIn    time.Time `json:"expireIn"`
}
