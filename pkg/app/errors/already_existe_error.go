package errors

type AlreadyExisteError struct {
	Message string
}

func (e AlreadyExisteError) Error() string {
	return e.Message
}

func NewEAlreadyExisteError(m string) AlreadyExisteError {
	return AlreadyExisteError{Message: m}
}
