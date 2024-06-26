// github.com/custom-package/errors
package errors

import "fmt"

type CustomError struct {
	error        // cause of the error
	msg   string // wrapped message
}

var DefaultSeparator string = "\n"

func New(msg string) error {
	return &CustomError{msg: msg}
}

func NewWithError(err error) error {
	return &CustomError{msg: err.Error(), error: err}
}

func WrapMessage(err error, msg string) error {
	if err == nil {
		return nil
	}
	return &CustomError{msg: msg, error: err}
}

func WrapMessageWithCode(err error, msg string, code string) error {
	if err == nil {
		return nil
	}
	return &CustomError{msg: msg, error: err}
}

func (e *CustomError) String() string {
	if e == nil {
		return ""
	}
	if e.error == nil {
		return e.msg
	}
	if stringer, ok := e.error.(fmt.Stringer); ok {
		return e.msg + DefaultSeparator + stringer.String()
	}
	return e.msg + DefaultSeparator + e.error.Error()
}

func (e *CustomError) Error() string { return e.msg }

func (e *CustomError) Cause() error { return e.error }

type causer interface {
	Cause() error
}

// The function "Cause" recursively retrieves the root cause of an error by checking if the error
// implements the "causer" interface.
func Cause(err error) error {
	if causer, ok := err.(causer); ok && causer.Cause() != nil {
		return Cause(causer.Cause())
	}
	return err
}

// The UnWrap function takes an error and return unwraped error.
func UnWrap(err error) error {
	if causer, ok := err.(causer); ok && causer != nil {
		return causer.Cause()
	}
	return err
}
