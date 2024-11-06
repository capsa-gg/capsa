package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"

	"github.com/lucianonooijen/capsa/server/internal/interactor"
	"github.com/lucianonooijen/capsa/server/internal/server/bodies"
)

// Extracts the response body and validates the data structure (for required fields).
// Returns the extracted data (as *T) and an error if one occurred.
// c.Abort is called if the data cannot be extracted or if the post body is invalid.
// Error handling is done in this helper, if err != nil, the calling code should just stop execution by calling return.
func extractBodyJSON[T any](c *gin.Context, s *interactor.Services) (*T, error) {
	log := s.Config.RootLogger.Named("extractBodyJSON").Sugar().
		With("method", c.Request.Method).
		With("uri", c.Request.RequestURI)

	var data T

	// Bind the request body
	if err := binding.JSON.Bind(c.Request, &data); err != nil {
		log.Infof("error binding JSON body: %s", err)

		res := bodies.ErrorResponse{Error: "invalid request body, cannot extract payload"}

		if s.Config.IsDevMode {
			res.RawError = err.Error()
		}

		c.JSON(http.StatusBadRequest, res)
		c.Abort()

		return nil, err
	}

	// Validate the body with `validate` struct tags
	validate := validator.New()
	if err := validate.Struct(data); err != nil {
		validationErr := createValidationError(err)

		res := bodies.ErrorResponse{
			Error:   "post body validation failed",
			Details: validationErr.Error(),
		}

		if s.Config.IsDevMode {
			res.RawError = err.Error()
		}

		log.With("raw_err", err).Infof("error validating post body: %s", validationErr.Error())

		c.JSON(http.StatusBadRequest, res)
		c.Abort()

		return nil, err
	}

	return &data, nil
}

// Creates a user-friendly validation error with the missing fields, based on validator.ValidationErrors.
func createValidationError(err error) error {
	// Convert to ValidationErrors type
	var validationErrors validator.ValidationErrors

	ok := errors.As(err, &validationErrors)
	if !ok {
		return err
	}

	// Build array with incorrect error fields and create error string
	var incorrectFields []string
	for _, e := range validationErrors {
		incorrectFields = append(incorrectFields, fmt.Sprintf("%s (%s)", e.Field(), e.Tag()))
	}

	incorrectFieldsString := strings.Join(incorrectFields, ", ")
	formattedError := fmt.Errorf("incorrect or missing fields in body: %s", incorrectFieldsString)

	// Return formatted error string
	return formattedError
}
