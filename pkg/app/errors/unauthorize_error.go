package errors

type UnauthorizeError struct {
	Message string
}

func (e UnauthorizeError) Error() string {
	return e.Message
}

func NewUnauthorizeError(m string) UnauthorizeError {
	return UnauthorizeError{Message: m}
}
