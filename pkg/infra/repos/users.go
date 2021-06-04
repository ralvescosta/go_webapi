package repos

import (
	"context"
	"database/sql"

	log "github.com/sirupsen/logrus"

	"webapi/pkg/app/dtos"
	"webapi/pkg/app/entities"
	"webapi/pkg/app/interfaces"
)

type userRepository struct {
	db *sql.DB
}

func (r userRepository) Create(ctx context.Context, user *dtos.UserDto) (*entities.User, error) {

	sql := "INSERT INTO users (first_name, last_name, email, password) VALUES ($1, $2, $3, $4) RETURNING *"

	prepare, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		log.Errorf("userRepository.Create - prepare statement: %v", err)
		return nil, err
	}

	entity := &entities.User{}
	err = prepare.QueryRowContext(
		ctx,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Password,
	).Scan(
		&entity.Id,
		&entity.FirstName,
		&entity.LastName,
		&entity.Email,
		&entity.Password,
		&entity.CreatedAt,
		&entity.UpdatedAt,
		&entity.DeletedAt,
	)
	if err != nil {
		log.Errorf("userRepository.Crete - query statement: %v", err)
		return nil, err
	}

	return entity, nil
}

func (r userRepository) FindByEmail(ctx context.Context, email string) (*entities.User, error) {
	sql :=
		`SELECT 
			id AS Id,
			first_name AS FirstName,
			last_name AS LastName,
			email AS Email,
			password as Password,
			created_at AS CreatedAt,
			updated_at AS UpdatedAt,
			deleted_at AS DeletedAt
		FROM users
		WHERE email = $1
	`
	prepare, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		log.Errorf("userRepository.FindByEmail - prepare statement: %v", err)
		return nil, err
	}

	entity := &entities.User{}

	err = prepare.QueryRowContext(ctx, &email).Scan(
		&entity.Id,
		&entity.FirstName,
		&entity.LastName,
		&entity.Email,
		&entity.Password,
		&entity.CreatedAt,
		&entity.UpdatedAt,
		&entity.DeletedAt,
	)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, nil
		}

		log.Errorf("userRepository.FindByEmail - query statement: %v", err)
		return nil, err
	}

	return entity, nil
}

func NewUserRepository(db *sql.DB) interfaces.IUserRepository {
	return &userRepository{db: db}
}
