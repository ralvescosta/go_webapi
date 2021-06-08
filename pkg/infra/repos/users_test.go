package repos

import (
	"context"
	"database/sql"
	"log"
	"testing"
	"time"
	"webapi/pkg/app/dtos"
	"webapi/pkg/app/entities"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func newDbMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

var user = &dtos.UserDto{
	FirstName: "name",
	LastName:  "name",
	Email:     "email@email.com",
	Password:  "123456",
}

func TestUserRepo_ShouldCreateUserCorrectly(t *testing.T) {
	db, mock := newDbMock()
	userRepo := NewUserRepository(db)

	query := "INSERT INTO users \\(first_name, last_name, email, password\\) VALUES \\(\\$1, \\$2, \\$3, \\$4\\) RETURNING \\*"
	rows := sqlmock.NewRows([]string{"id", "first_name", "last_name", "email", "password", "created_at", "updated_at", "deleted_at"}).
		AddRow(1, user.FirstName, user.LastName, user.Email, user.Password, time.Now(), time.Now(), nil)

	prep := mock.ExpectPrepare(query)
	prep.ExpectQuery().WithArgs(user.FirstName, user.LastName, user.Email, user.Password).WillReturnRows(rows)

	result, err := userRepo.Create(context.Background(), user)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.IsType(t, result, &entities.User{})
	assert.Equal(t, result.Email, user.Email)
}

func TestUserRepo_ShouldOccurErrInPrepareStatement(t *testing.T) {
	db, mock := newDbMock()
	userRepo := NewUserRepository(db)

	mock.ExpectPrepare("")

	result, err := userRepo.Create(context.Background(), user)

	assert.NotNil(t, err)
	assert.Nil(t, result)
}

func TestUserRepo_ShouldOccurErrBeforeScanStatatment(t *testing.T) {
	db, mock := newDbMock()
	userRepo := NewUserRepository(db)

	query := "INSERT INTO users \\(first_name, last_name, email, password\\) VALUES \\(\\$1, \\$2, \\$3, \\$4\\) RETURNING \\*"
	rows := sqlmock.NewRows([]string{"id", "first_name", "last_name", "email", "password", "created_at", "updated_at", "deleted_at"}).
		AddRow(1, user.FirstName, user.LastName, user.Email, user.Password, time.Now().String(), time.Now().String(), nil)

	prep := mock.ExpectPrepare(query)
	prep.ExpectQuery().WithArgs(user.FirstName, user.LastName, user.Email, user.Password).WillReturnRows(rows)

	result, err := userRepo.Create(context.Background(), user)

	assert.NotNil(t, err)
	assert.Nil(t, result)
}
