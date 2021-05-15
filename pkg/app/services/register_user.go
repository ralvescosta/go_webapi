package services

import (
	"webapi/pkg/app/dtos"
	"webapi/pkg/app/interfaces"
)

type IUserService interface {
	Register(user *dtos.UserDto) error
}

type userService struct {
	repo interfaces.IUserRepository
}

func (s *userService) Register(user *dtos.UserDto) error {
	return nil
}

func NewUserService() IUserService {
	return &userService{}
}
