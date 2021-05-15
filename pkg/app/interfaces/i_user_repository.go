package interfaces

import "webapi/pkg/app/dtos"

type IUserRepository interface {
	Create(user *dtos.UserDto)
}
