package mocks

import (
	"context"
	"webapi/pkg/app/dtos"
)

type UserServiceMocked struct{}

func (m UserServiceMocked) Register(ctx context.Context, user *dtos.UserDto) error {
	return nil
}

func NewUserServicesMock() *UserServiceMocked {
	return &UserServiceMocked{}
}
