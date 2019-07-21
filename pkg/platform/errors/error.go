package errors

import (
	"fmt"
)

// Application error codes
const (
	ECONFLICT     = "conflict"       // action cannot be performed
	EINTERNAL     = "internal"       // internal error
	EINVALID      = "invalid"        // validation failed
	ENOTFOUND     = "not_found"      // entity does not exist
	EUNAUTHORIZED = "not_authorized" //not authorized

	DefaultErrMessage = "An internal error has occurred"
)

// Error defines a standard application error
type Error struct {
	Code    string
	Message string
	Op      string
	Err     error
}

// New returns a pointer to an Error concrete
func New(c, m, op string, err error) *Error {
	return &Error{
		Code:    c,
		Message: m,
		Op:      op,
		Err:     err,
	}
}

// Error returns the string representation of the error message.
func (e *Error) Error() string {
	if e == nil {
		return ""
	}

	var details string
	if e.Err != nil {
		details = fmt.Sprintf("details: %s.", e.Err.Error())
	}

	return fmt.Sprintf("%s: [%s] %s. %s", e.Op, e.Code, e.Message, details)
}

// Code returns the code of the root error, if available. Otherwise returns EINTERNAL.
func Code(err error) string {
	if err == nil {
		return ""
	}

	if e, ok := err.(*Error); ok && e.Code != "" {
		return e.Code
	} else if ok && e.Err != nil {
		return Code(e.Err)
	}

	return EINTERNAL
}

// Message returns the human-readable message of the error, if available.
// Otherwise returns a generic error message.
func Message(err error) string {
	if err == nil {
		return ""
	}

	if e, ok := err.(*Error); ok && e.Message != "" {
		return e.Message
	} else if ok && e.Err != nil {
		return Message(e.Err)
	}

	return DefaultErrMessage
}

// Wrap adds message to error
func Wrap(err error, message string) error {
	if err == nil {
		return nil
	}

	if e, ok := err.(*Error); ok {
		if e != nil {
			e.Message = fmt.Sprintf("%s: %s", message, e.Message)
			return e
		}
	}

	return err
}
