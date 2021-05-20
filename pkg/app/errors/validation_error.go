package errors

type ValidationError struct {
	Message string
}

func (e ValidationError) Error() string {
	return e.Message
}

func NewValidationError(m string) ValidationError {
	return ValidationError{Message: m}
}
