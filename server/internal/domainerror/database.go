package domainerror

import (
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

// NewFromDatabaseError accepts a database error and returns a domain error.
func NewFromDatabaseError(err error) Error {
	// Not found in database => 404
	if errors.Is(err, pgx.ErrNoRows) {
		detail := fmt.Sprintf("Database not found error: %s", err)

		return Error{
			Type:     NotFound,
			Message:  "item not found",
			Details:  detail,
			RawError: err,
		}
	}

	var pqErr *pgconn.PgError
	ok := errors.As(err, &pqErr)

	if !ok {
		return New(Unexpected, "unexpected error", err)
	}

	if pqErr == nil {
		return New(Unexpected, "unexpected error", fmt.Errorf("pgErr is nil: %w", pqErr))
	}

	formatDetails := fmt.Errorf("%s (code %s, message %s)", pqErr.Detail, pqErr.Code, pqErr.Message)

	// Note: for finding error codes, see: https://www.postgresql.org/docs/11/errcodes-appendix.html or https://github.com/lib/pq/blob/master/error.go#L78
	switch pqErr.Code {
	case "23505": // unique_violation
		detail := "Data unique violation: " + pqErr.Detail

		return Error{
			Type:     Conflict,
			Message:  "database conflict",
			Details:  detail,
			RawError: err,
		}
	case "22P02": // invalid_text_representation
		detail := fmt.Sprintf("invalid input for enum type: %s (%s)", pqErr.Message, formatDetails)

		return Error{
			Type:     InvalidArgument,
			Message:  "incorrect enum member",
			Details:  detail,
			RawError: err,
		}
	case "23503": // foreign_key_violation
		detail := fmt.Sprintf("Foreign key database error: fkey = %s, message = %s", pqErr.ConstraintName, pqErr.Message)

		return Error{
			Type:     InvalidArgument,
			Message:  "foreign key violation",
			Details:  detail,
			RawError: err,
		}
	case "23514": // check_violation
		detail := fmt.Sprintf("Input does not satisfy database data check: check_name = %s, message = %s", pqErr.ConstraintName, pqErr.Message)

		return Error{
			Type:     InvalidArgument,
			Message:  "data constraints not met",
			Details:  detail,
			RawError: err,
		}
	default:
		return Error{
			Type:     Unexpected,
			Message:  "unexpected database error",
			RawError: err,
		}
	}
}
