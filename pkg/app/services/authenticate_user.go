package services

import "webapi/pkg/app/entities"

type IAuthenticationUser interface {
	Perform(email string, password string) (*entities.AuthenticatedUser, error)
}

type authenticationUser struct{}

func (s authenticationUser) Perform(email string, password string) (*entities.AuthenticatedUser, error) {
	return nil, nil
}

func NewAuthenticationUser() IAuthenticationUser {
	return &authenticationUser{}
}
