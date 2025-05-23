package errors

import "fmt"

type SqlError struct {
	SqlRequest string
	Message    string
	Err        error
}

func (e *SqlError) Error() string {
	return fmt.Sprintf("Error SQL :\n\t- SQL request : %s\n\t- Message : %s\n\t- Error : %o", e.SqlRequest, e.Message, e.Err)
}

func NewSqlError(sqlRequest string, message string, err error) *SqlError {
	return &SqlError{
		SqlRequest: sqlRequest,
		Message: message,
		Err: err,
	}
}