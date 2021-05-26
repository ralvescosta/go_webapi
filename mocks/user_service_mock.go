package mocks

import (
	"webapi/pkg/app/dtos"
)

type UserServiceMocked struct{}

func (m UserServiceMocked) Register(user *dtos.UserDto) error {
	return nil
}

func NewUserServicesMock() *UserServiceMocked {
	return &UserServiceMocked{}
}
