package errors

import "fmt"

type HtmlError struct {
	Code    int
	Message string
	Err     error
}

func (e *HtmlError) Error() string {
	return fmt.Sprintf("%d - %s", e.Code, e.Message)
}

func NewHtmlError(code int, message string, err error) *HtmlError {
	return &HtmlError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}
