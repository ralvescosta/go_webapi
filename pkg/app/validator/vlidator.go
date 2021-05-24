package validator

import (
	"fmt"
	"regexp"

	err "webapi/pkg/app/errors"
)

type Validator struct {
	value        interface{}
	propertyName string
	Messages     []string
	IsValid      bool
}

func (v *Validator) Set(value interface{}, propertyName string) *Validator {
	v.value = value
	v.propertyName = propertyName
	return v
}

func (v *Validator) Email() *Validator {
	toString, ok := v.value.(string)
	if !ok {
		v.Messages = append(v.Messages, fmt.Sprintf("Property: %s - Email in Wrong Format", v.propertyName))
		return v
	}

	re := regexp.MustCompile("(?:[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*|\"(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21\x23-\x5b\x5d-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])*\")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\\[(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?|[a-z0-9-]*[a-z0-9]:(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21-\x5a\x53-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])+)\\])")
	suc := re.MatchString(fmt.Sprintf("%v", toString))
	if !suc {
		v.Messages = append(v.Messages, fmt.Sprintf("Property: %s - Email in Wrong Format", v.propertyName))
	}

	return v
}

func (v *Validator) Required() *Validator {
	if v.value == nil || v.value == "" || v.value == 0 || v.value == 0.0 {
		v.Messages = append(v.Messages, fmt.Sprintf("Property: %s - Is Required", v.propertyName))
	}
	return v
}

func (v *Validator) MinInt(min int) *Validator {
	minInt, ok := v.value.(int)
	if !ok {
		v.Messages = append(v.Messages, fmt.Sprintf("Property: %s - Invalid Type", v.propertyName))
		return v
	}

	if minInt < min {
		v.Messages = append(v.Messages, fmt.Sprintf("Property: %s - Is Required", v.propertyName))
	}
	return v
}

func (v *Validator) MinString(min int) *Validator {
	toString, ok := v.value.(string)
	if !ok {
		v.Messages = append(v.Messages, fmt.Sprintf("Property: %s - Invalid Type", v.propertyName))
		return v
	}

	if len(toString) < min {
		v.Messages = append(v.Messages, fmt.Sprintf("Property: %s - Is Required", v.propertyName))
	}
	return v
}

func (v *Validator) Validate() error {
	if len(v.Messages) > 0 {
		v.IsValid = false
		return err.NewValidationError("Invalide")
	}

	return nil
}
