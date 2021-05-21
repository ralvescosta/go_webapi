package mocks

import (
	"webapi/pkg/app/dtos"
	"webapi/pkg/app/entities"
	"webapi/pkg/app/interfaces"
)

type userRepositoryInMemory struct{}

func (m userRepositoryInMemory) Create(user *dtos.UserDto) (*entities.User, error) {
	return nil, nil
}

func (m userRepositoryInMemory) FindByEmail(email string) (*entities.User, error) {
	return nil, nil
}

func NewUserRepositoryInMemory() interfaces.IUserRepository {
	return &userRepositoryInMemory{}
}
