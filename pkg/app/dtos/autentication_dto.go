package dtos

import (
	"regexp"
	err "webapi/pkg/app/errors"
)

type AuthenticationDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (a AuthenticationDTO) Validate() error {
	if a.Email == "" {
		return err.NewValidationError("email is required")
	}
	if a.Password == "" {
		return err.NewValidationError("email is required")
	}
	if len(a.Password) < 6 {
		return err.NewValidationError("Password to small")
	}

	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	suc := re.MatchString(a.Email)
	if !suc {
		return err.NewValidationError("Email in Wrong Format")
	}

	return nil
}
