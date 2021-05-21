package services

import (
	"time"
	"webapi/pkg/app/dtos"
	"webapi/pkg/app/entities"
	"webapi/pkg/app/errors"
	i "webapi/pkg/app/interfaces"
)

type IAuthenticationUser interface {
	Perform(email string, password string) (*entities.AuthenticatedUser, error)
}

type authenticationUser struct {
	repo         i.IUserRepository
	hash         i.IHasher
	tokenManager i.ITokenManager
}

func (s authenticationUser) Perform(email string, password string) (*entities.AuthenticatedUser, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return nil, errors.NewInternalError(err.Error())
	}

	if user == nil || user.Email == "" || user.Password == "" {
		errors.NewBadRequestError("Email is wrong or not exist")
	}
	if isPasswordValid := s.hash.Verify(password, user.Password); !isPasswordValid {
		return nil, errors.NewBadRequestError("User Password is wrong")
	}

	tokenDataDto := &dtos.TokenDataDto{
		Id:       user.Id,
		ExpireIn: time.Now().Add(time.Hour),
		Audience: "1.1.1",
	}
	accessToken, err := s.tokenManager.GenerateToken(tokenDataDto)
	if err != nil {
		return nil, errors.NewInternalError("An Error occur while processing the request, try again!")
	}

	return &entities.AuthenticatedUser{
		Id:          user.Id,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		AccessToken: accessToken,
		ExpireIn:    tokenDataDto.ExpireIn,
	}, nil
}

func NewAuthenticationUser(repo i.IUserRepository, hash i.IHasher, tokenManager i.ITokenManager) IAuthenticationUser {
	return &authenticationUser{repo: repo, hash: hash, tokenManager: tokenManager}
}
