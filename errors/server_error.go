package errors

type ServerError struct {
	Message string
	Err     error
}

func (e *ServerError) Error() string {
	return e.Message
}

func NewServerError(message string, err error) error {
	return &ServerError{
		Message: message,
		Err:     err,
	}
}
