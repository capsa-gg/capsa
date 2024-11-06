package entities

import (
	"errors"
	"fmt"

	"github.com/lib/pq"
)

// DomainErrorType is used to differentiate between different types of domain errors.
type DomainErrorType string

const (
	// DomainErrorInvalidArgument indicates that validation on the arguments passed in has failed.
	DomainErrorInvalidArgument = DomainErrorType("error-invalid-argument")

	// DomainErrorNoPermission indicates that requester does not have permission to access the requested data .
	DomainErrorNoPermission = DomainErrorType("error-invalid-no-permission")

	// DomainErrorNotFound indicates that a requested resource was not found.
	DomainErrorNotFound = DomainErrorType("error-not-found")

	// DomainErrorConflict indicates that a data conflict occurred.
	DomainErrorConflict = DomainErrorType("error-conflict")

	// DomainErrorUnexpected indicates that an unexpected error occurred.
	DomainErrorUnexpected = DomainErrorType("error-unexpected")
)

// DomainError is used by domain logic to return errors to the calling code.
type DomainError struct {
	Type     DomainErrorType
	Message  string
	Details  string
	RawError error
}

// Error implements the error interface.
func (de DomainError) Error() string {
	return fmt.Sprintf("[]: %s", de.Message)
}

// NewDomainError returns a new instance of DomainError.
func NewDomainError(errorType DomainErrorType, message string, err error) DomainError {
	return DomainError{
		Type:     errorType,
		Message:  message,
		RawError: err,
	}
}

// NewDomainErrorFromDatabaseError accepts a database error and returns a domain error.
func NewDomainErrorFromDatabaseError(err error) DomainError {
	var pqErr *pq.Error
	ok := errors.As(err, &pqErr)

	if !ok {
		return NewDomainError(DomainErrorUnexpected, "unexpected error", err)
	}

	formatDetails := fmt.Errorf("%s (code %s, name %s)", pqErr.Detail, pqErr.Code, pqErr.Code.Name())

	switch pqErr.Code.Name() {
	case "unique_violation":
		detail := fmt.Sprintf("Data unique violation: %s", pqErr.Detail)

		return DomainError{
			Type:     DomainErrorConflict,
			Message:  "database conflict",
			Details:  detail,
			RawError: err,
		}
	case "invalid_text_representation":
		detail := fmt.Sprintf("invalid input for enum type: %s (%s)", pqErr.Message, formatDetails)

		return DomainError{
			Type:     DomainErrorInvalidArgument,
			Message:  "incorrect enum member",
			Details:  detail,
			RawError: err,
		}
	case "foreign_key_violation":
		detail := fmt.Sprintf("Foreign key database error: fkey = %s, message = %s", pqErr.Constraint, pqErr.Message)

		return DomainError{
			Type:     DomainErrorInvalidArgument,
			Message:  "foreign key violation",
			Details:  detail,
			RawError: err,
		}
	case "check_violation":
		detail := fmt.Sprintf("Input does not satisfy database data check: check_name = %s, message = %s", pqErr.Constraint, pqErr.Message)

		return DomainError{
			Type:     DomainErrorInvalidArgument,
			Message:  "data constraints not met",
			Details:  detail,
			RawError: err,
		}
	default:
		return DomainError{
			Type:     DomainErrorUnexpected,
			Message:  "unexpected database error",
			RawError: err,
		}
	}
}
