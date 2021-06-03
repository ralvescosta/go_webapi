package mocks

import (
	"context"
	"webapi/pkg/app/entities"
)

type AuthenticationUserMocked struct{}

func (m AuthenticationUserMocked) Perform(ctx context.Context, email, password, audience string) (*entities.AuthenticatedUser, error) {
	return nil, nil
}

func NewAuthenticationUserMocked() *AuthenticationUserMocked {
	return &AuthenticationUserMocked{}
}
