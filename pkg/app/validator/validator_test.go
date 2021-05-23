package validator

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSet_ShouldExecSetCorrectly(t *testing.T) {
	var validator Validator = Validator{}
	value := "some value"
	propertyName := "SomeProperty"

	validator.Set(value, propertyName)

	assert.Equal(t, validator.value, value)
	assert.Equal(t, validator.propertyName, propertyName)
}

func TestRequired_ShouldDontAddMessageIfThereIsValue(t *testing.T) {
	var validator Validator = Validator{}
	value := "some value"
	propertyName := "SomeProperty"

	validator.Set(value, propertyName).Required()

	assert.Empty(t, validator.Messages)
}

func TestRequired_ShouldAddMessageIfThereIsNoValue(t *testing.T) {
	var validator Validator = Validator{}
	value := ""
	propertyName := "SomeProperty"

	validator.Set(value, propertyName).Required()

	assert.Equal(t, validator.Messages, []string{fmt.Sprintf("Property: %s - Is Required", propertyName)})
}

func TestMinInt_ShouldDontAddMessageIfValueIsBiggerThen(t *testing.T) {
	var validator Validator = Validator{}
	value := 5
	propertyName := "SomeProperty"

	validator.Set(value, propertyName).MinInt(1)

	assert.Empty(t, validator.Messages)
}

func TestMinInt_ShouldAddMessageIfValueIsNoBiggerThen(t *testing.T) {
	var validator Validator = Validator{}
	value := 5
	propertyName := "SomeProperty"

	validator.Set(value, propertyName).MinInt(6)

	assert.Equal(t, validator.Messages, []string{fmt.Sprintf("Property: %s - Is Required", propertyName)})
}

func TestMinString_ShouldDontAddMessageIfValueIsBiggerThen(t *testing.T) {
	var validator Validator = Validator{}
	value := "something"
	propertyName := "SomeProperty"

	validator.Set(value, propertyName).MinString(1)

	assert.Empty(t, validator.Messages)
}

func TestMinString_ShouldAddMessageIfValueIsNoBiggerThen(t *testing.T) {
	var validator Validator = Validator{}
	value := "something"
	propertyName := "SomeProperty"

	validator.Set(value, propertyName).MinString(20)

	assert.Equal(t, validator.Messages, []string{fmt.Sprintf("Property: %s - Is Required", propertyName)})
}

func TestEmail_ShouldDontAddMessageIfValueIsValidEmail(t *testing.T) {
	var validator Validator = Validator{}
	value := "email@email.com"
	propertyName := "SomeProperty"

	validator.Set(value, propertyName).Email()

	assert.Empty(t, validator.Messages)
}

func TestEmail_ShouldAddMessageIfValueIsNoValidEmail(t *testing.T) {
	var validator Validator = Validator{}
	value := "wrong@email"
	propertyName := "SomeProperty"

	validator.Set(value, propertyName).Email()

	assert.Equal(t, validator.Messages, []string{fmt.Sprintf("Property: %s - Email in Wrong Format", propertyName)})
}

func TestValidate_ShouldReturnNilIfNoErrorHappened(t *testing.T) {
	var validator Validator = Validator{}
	value := "something"
	propertyName := "SomeProperty"

	validator.Set(value, propertyName).MinString(1)
	err := validator.Validate()

	assert.Nil(t, err)
}

func TestValidate_ShouldReturnValidationErrorIfSomeErrorHappened(t *testing.T) {
	var validator Validator = Validator{}
	value := "something"
	propertyName := "SomeProperty"

	validator.Set(value, propertyName).MinString(30)
	err := validator.Validate()

	assert.Error(t, err)
}
