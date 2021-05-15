package errors

type InternalError struct {
	Message string
}

func (e InternalError) Error() string {
	return e.Message
}

func NewInternalError(m string) InternalError {
	return InternalError{Message: m}
}
