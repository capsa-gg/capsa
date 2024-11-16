package entities

import (
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
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
	return fmt.Sprintf("[%s]: %s", de.Type, de.Message)
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
	// Not found in database => 404
	if errors.Is(err, pgx.ErrNoRows) {
		detail := fmt.Sprintf("Database not found error: %s", err)

		return DomainError{
			Type:     DomainErrorNotFound,
			Message:  "item not found",
			Details:  detail,
			RawError: err,
		}
	}

	var pqErr *pgconn.PgError
	ok := errors.As(err, &pqErr)

	if !ok {
		return NewDomainError(DomainErrorUnexpected, "unexpected error", err)
	}

	formatDetails := fmt.Errorf("%s (code %s, message %s)", pqErr.Detail, pqErr.Code, pqErr.Message)

	// Note: for finding error codes, see: https://www.postgresql.org/docs/11/errcodes-appendix.html or https://github.com/lib/pq/blob/master/error.go#L78
	switch pqErr.Code {
	case "23505": // unique_violation
		detail := fmt.Sprintf("Data unique violation: %s", pqErr.Detail)

		return DomainError{
			Type:     DomainErrorConflict,
			Message:  "database conflict",
			Details:  detail,
			RawError: err,
		}
	case "22P02": // invalid_text_representation
		detail := fmt.Sprintf("invalid input for enum type: %s (%s)", pqErr.Message, formatDetails)

		return DomainError{
			Type:     DomainErrorInvalidArgument,
			Message:  "incorrect enum member",
			Details:  detail,
			RawError: err,
		}
	case "23503": // foreign_key_violation
		detail := fmt.Sprintf("Foreign key database error: fkey = %s, message = %s", pqErr.ConstraintName, pqErr.Message)

		return DomainError{
			Type:     DomainErrorInvalidArgument,
			Message:  "foreign key violation",
			Details:  detail,
			RawError: err,
		}
	case "23514": // check_violation
		detail := fmt.Sprintf("Input does not satisfy database data check: check_name = %s, message = %s", pqErr.ConstraintName, pqErr.Message)

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
