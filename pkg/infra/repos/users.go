package repos

import (
	"database/sql"

	"webapi/pkg/app/dtos"
	"webapi/pkg/app/entities"
	"webapi/pkg/app/interfaces"
)

type userRepository struct {
	db *sql.DB
}

func (r userRepository) Create(user *dtos.UserDto) (entities.User, error) {
	return entities.User{}, nil
}

func NewUserRepository(db *sql.DB) interfaces.IUserRepository {
	return &userRepository{db: db}
}
