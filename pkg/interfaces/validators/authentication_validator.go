package validators

import (
	"webapi/pkg/app/dtos"
	"webapi/pkg/app/validator"
)

func ValidateAuthUserBody(dto *dtos.AuthenticationDTO) ([]string, error) {
	validator := validator.Validator{}

	validator.Set(dto.Email, "Email").Email().Required()
	validator.Set(dto.Password, "Password").MinString(6).Required()

	return validator.Messages, validator.Validate()
}
