package services

import (
	"webapi/pkg/app/entities"
	"webapi/pkg/app/errors"
	i "webapi/pkg/app/interfaces"
)

type IAuthenticationUser interface {
	Perform(email string, password string) (*entities.AuthenticatedUser, error)
}

type authenticationUser struct {
	repo i.IUserRepository
	hash i.IHasher
}

func (s authenticationUser) Perform(email string, password string) (*entities.AuthenticatedUser, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	if isPasswordValid := s.hash.Verify(email, user.Password); !isPasswordValid {
		return nil, errors.NewUnauthorizeError("User Unauthorized")
	}

	return nil, nil
}

func NewAuthenticationUser(repo i.IUserRepository, hash i.IHasher) IAuthenticationUser {
	return &authenticationUser{repo: repo, hash: hash}
}
