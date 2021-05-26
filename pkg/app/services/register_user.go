package services

import (
	"webapi/pkg/app/dtos"
	i "webapi/pkg/app/interfaces"
)

type IUserService interface {
	Register(user *dtos.UserDto) error
}

type userService struct {
	repo i.IUserRepository
	hash i.IHasher
}

func (s *userService) Register(user *dtos.UserDto) error {
	passHashed, err := s.hash.Hahser(user.Password)
	if err != nil {
		return err
	}

	user.Password = passHashed
	_, err = s.repo.Create(user)
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
