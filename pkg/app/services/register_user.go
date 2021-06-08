package services

import (
	"context"
	"webapi/pkg/app/dtos"
	"webapi/pkg/app/errors"
	i "webapi/pkg/app/interfaces"
)

type IUserService interface {
	Register(ctx context.Context, user *dtos.UserDto) error
}

type userService struct {
	repo i.IUserRepository
	hash i.IHasher
}

func (s *userService) Register(ctx context.Context, user *dtos.UserDto) error {
	alreadyExist, err := s.repo.FindByEmail(ctx, user.Email)
	if err != nil {
		return err
	}
	if alreadyExist != nil && alreadyExist.Email != "" {
		return errors.NewAlreadyExisteError("user already exist")
	}

	passHashed, err := s.hash.Hahser(user.Password)
	if err != nil {
		return err
	}

	user.Password = passHashed
	_, err = s.repo.Create(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func NewUserService(repo i.IUserRepository, hash i.IHasher) IUserService {
	return &userService{
		repo: repo,
		hash: hash,
	}
}
