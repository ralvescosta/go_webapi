package errors

type NotFoundError struct {
	Message string
}

func (e NotFoundError) Error() string {
	return e.Message
}

func NewNotFoundError(m string) NotFoundError {
	return NotFoundError{Message: m}
}
