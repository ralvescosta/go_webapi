package dtos

import (
	"regexp"

	err "webapi/pkg/app/errors"
)

type UserDto struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (u UserDto) Validate() error {
	if u.FirstName == "" {
		return err.NewEntityError("fistName is required")
	}
	if u.LastName == "" {
		return err.NewEntityError("lastName is required")
	}
	if u.Email == "" {
		return err.NewEntityError("email is required")
	}
	if u.Password == "" {
		return err.NewEntityError("pasword is required")
	}

	if len(u.Password) < 6 {
		return err.NewEntityError("Password to small")
	}

	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	suc := re.MatchString(u.Email)
	if !suc {
		return err.NewEntityError("Email in Wrong Format")
	}

	return nil
}
