package mocks

import "webapi/pkg/app/entities"

type AuthenticationUserMocked struct{}

func (m AuthenticationUserMocked) Perform(email, password, audience string) (*entities.AuthenticatedUser, error) {
	return nil, nil
}

func NewAuthenticationUserMocked() *AuthenticationUserMocked {
	return &AuthenticationUserMocked{}
}
