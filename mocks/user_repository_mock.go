package mocks

import (
	"errors"
	"webapi/pkg/app/dtos"
	"webapi/pkg/app/entities"
	"webapi/pkg/app/interfaces"
)

type userRepositoryInMemory struct {
	failure          bool
	mockedUserEntity *entities.User
}

func (m userRepositoryInMemory) Create(user *dtos.UserDto) (*entities.User, error) {
	if m.failure {
		return nil, errors.New("Error")
	}
	return m.mockedUserEntity, nil
}

func (m userRepositoryInMemory) FindByEmail(email string) (*entities.User, error) {
	if m.failure {
		return nil, errors.New("Error")
	}
	return m.mockedUserEntity, nil
}

func NewUserRepositoryInMemory(failure bool, mockedUserEntity *entities.User) interfaces.IUserRepository {
	return &userRepositoryInMemory{failure: failure, mockedUserEntity: mockedUserEntity}
}
