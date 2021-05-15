package interfaces

import (
	"webapi/pkg/app/dtos"
	"webapi/pkg/app/entities"
)

type IUserRepository interface {
	Create(user *dtos.UserDto) (entities.User, error)
}
