package mocks

import (
	"webapi/pkg/app/dtos"
)

type userServiceMocked struct{}

func (m userServiceMocked) Register(user *dtos.UserDto) error {
	return nil
}

func NewUserServicesMock() *userServiceMocked {
	return &userServiceMocked{}
}
