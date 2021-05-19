package entities

import "time"

type AuthenticatedUser struct {
	Id          int
	FirstName   string
	LastName    string
	Email       string
	AccessToken string
	ExpireIn    time.Time
}
