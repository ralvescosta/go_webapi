package errors

type BadRequestError struct {
	Message string
}

func (e BadRequestError) Error() string {
	return e.Message
}

func NewBadRequestError(m string) BadRequestError {
	return BadRequestError{Message: m}
}
