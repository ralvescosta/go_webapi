package validators

import (
	"webapi/pkg/app/dtos"
	"webapi/pkg/app/validator"
)

func ValidateCreateUserBody(data *dtos.UserDto) ([]string, error) {
	validator := validator.Validator{}

	validator.Set(data.FirstName, "FirstName").Required()
	validator.Set(data.LastName, "LastName").Required()
	validator.Set(data.Email, "Email").Email().Required()
	validator.Set(data.Password, "Password").MinString(6).Required()

	return validator.Messages, validator.Validate()
}
