package domainerror

import (
	"fmt"
)

// Error is used by domain logic to return errors to the calling code.
type Error struct {
	Type     ErrorType
	Message  string
	Details  string
	RawError error
}

// Error implements the error interface.
func (de Error) Error() string {
	return fmt.Sprintf("[%s]: %s", de.Type, de.Message)
}

// New returns a new instance of Error.
func New(errorType ErrorType, message string, err error) Error {
	return Error{
		Type:     errorType,
		Message:  message,
		RawError: err,
	}
}
