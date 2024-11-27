package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/capsa-gg/capsa/server/internal/domain/user"
	"github.com/capsa-gg/capsa/server/internal/server/bodies"
)

// UserPasswordRequest allows users to request a password reset
// @Summary 	User password reset request
// @Tags        UserAuthentication
// @Produce     json
// @Param		email 	query		string 		true 	"UserEmail"
// @Description	Allows clients to create a log session and receive a token to send logs with
// @Success		201
// @Failure     400		{object}	bodies.ErrorResponse
// @Failure     404		{object}	bodies.ErrorResponse
// @Failure     500		{object}	bodies.ErrorResponse
// @Header		all		{string} 	X-Capsa-Server-Version		"Current Capsa Server version"
// @Router 		/user/auth/password-reset [get]
func (h Handlers) UserPasswordRequest(c *gin.Context) {
	log := h.logger.Named("UserLogin")

	email := c.Query("email")
	if email == "" {
		c.JSON(http.StatusBadRequest, bodies.ErrorResponse{Error: "email field is required"})

		return
	}

	log = log.With("email", email)
	log.Debug("attempting to start password reset flow")

	err := user.PasswordResetStart(c, h.services, email)
	if err != nil {
		h.sendErrorResponse(c, err)

		return
	}

	log.Info("password reset initialized")

	c.Status(http.StatusCreated)
}

// UserPasswordComplete allows users to complete a password reset
// @Summary 	User password reset set password
// @Tags        UserAuthentication
// @Produce     json
// @Param		email 					query	string 										true 	"UserEmail"
// @Param		password_reset_data 	body 	bodies.UserPasswordResetCompleteRequest 	true 	"UserPasswordCompleteRequest"
// @Description	Allows clients to create a log session and receive a token to send logs with
// @Success		201
// @Failure     400		{object}	bodies.ErrorResponse
// @Failure     404		{object}	bodies.ErrorResponse
// @Failure     500		{object}	bodies.ErrorResponse
// @Header		all		{string} 	X-Capsa-Server-Version		"Current Capsa Server version"
// @Router 		/user/auth/password-reset [post]
func (h Handlers) UserPasswordComplete(c *gin.Context) {
	log := h.logger.Named("UserLogin")

	req, err := extractBodyJSON[bodies.UserPasswordResetCompleteRequest](c, h.services)
	if err != nil {
		return // Error sent by extractBodyJSON
	}

	if req == nil {
		h.sendErrorResponse(c, jsonBodyExtractionError)

		return
	}

	log = log.With("email", req.ResetToken)
	log.Debug("attempting to complete password reset flow")

	err = user.PasswordResetComplete(c, h.services, req.ResetToken, req.Password)
	if err != nil {
		h.sendErrorResponse(c, err)

		return
	}

	log.Info("password reset completed")

	c.Status(http.StatusCreated)
}
