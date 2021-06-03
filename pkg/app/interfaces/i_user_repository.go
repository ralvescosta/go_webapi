package interfaces

import (
	"context"
	"webapi/pkg/app/dtos"
	"webapi/pkg/app/entities"
)

type IUserRepository interface {
	Create(ctx context.Context, user *dtos.UserDto) (*entities.User, error)
	FindByEmail(ctx context.Context, email string) (*entities.User, error)
}
