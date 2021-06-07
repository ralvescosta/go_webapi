package repos

import (
	"context"
	"database/sql"
	"log"
	"testing"
	"time"
	"webapi/pkg/app/dtos"

	"github.com/DATA-DOG/go-sqlmock"
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

func TestT(t *testing.T) {
	db, mock := newDbMock()
	userRepo := NewUserRepository(db)

	query := "INSERT INTO users \\(first_name, last_name, email, password\\) VALUES \\(\\?, \\?, \\?, \\?\\) RETURNING \\*"
	rows := sqlmock.NewRows([]string{"id", "first_name", "last_name", "email", "password", "created_at", "updated_at", "deleted_at"}).
		AddRow(1, user.FirstName, user.LastName, user.Email, user.Password, time.Now().Format("2006-01-02 15:04:05"), time.Now().Format("2006-01-02 15:04:05"), nil)

	prep := mock.ExpectPrepare(query)
	prep.ExpectQuery().WithArgs(user.FirstName, user.LastName, user.Email, user.Password).WillReturnRows(rows)

	userRepo.Create(context.Background(), user)

	// fmt.Print(result, err)
}
