package validators

import (
	"webapi/pkg/app/dtos"
	"webapi/pkg/app/validator"
)

func ValidateCreateUserBody(data *dtos.UserDto) (error, []string) {
	validator := validator.Validator{}

	validator.Set(data.FirstName, "FirstName").Required()
	validator.Set(data.LastName, "LastName").Required()
	validator.Set(data.Email, "Email").Email().Required()
	validator.Set(data.Password, "Password").MinString(6).Required()

	return validator.Validate(), validator.Messages
}
