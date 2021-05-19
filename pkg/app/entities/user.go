package entities

import (
	"time"
)

type User struct {
	Id        uint32
	FirstName string
	LastName  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
