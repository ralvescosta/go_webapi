package errors

type EntityError struct {
	Message string
}

func (e EntityError) Error() string {
	return e.Message
}

func NewEntityError(m string) EntityError {
	return EntityError{Message: m}
}
