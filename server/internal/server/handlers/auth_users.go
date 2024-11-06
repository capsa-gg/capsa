package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/lucianonooijen/capsa/server/internal/domain/user"
	"github.com/lucianonooijen/capsa/server/internal/server/bodies"
)

// UserLogin returns the user's information in case authentication was successful
// @Summary 	User login handler
// @Tags        UserAuthentication
// @Accept 		json
// @Produce     json
// @Param		login_request 	body 	bodies.UserLoginRequest 	true 	"UserLoginRequest"
// @Description	Allows clients to create a log session and receive a token to send logs with
// @Success		200		{object}	entities.UserLoginResult
// @Failure     400		{object}	bodies.ErrorResponse
// @Failure     401		{object}	bodies.ErrorResponse
// @Failure     403		{object}	bodies.ErrorResponse
// @Failure     404		{object}	bodies.ErrorResponse
// @Failure     409		{object}	bodies.ErrorResponse
// @Failure     500		{object}	bodies.ErrorResponse
// @Header		all		{string} 	X-Capsa-Server-Version		"Current Capsa Server version"
// @Router 		/user/auth/login [post]
func (h Handlers) UserLogin(c *gin.Context) {
	log := h.logger.Named("UserLogin")

	req, err := extractBodyJSON[bodies.UserLoginRequest](c, h.services)
	if err != nil {
		return // Error sent by extractBodyJSON
	}

	log = log.With("user_email", req.Email)

	loginInfo, err := user.Login(h.services, req.Email, req.Password)
	if err != nil {
		log.Warn("failed login attempt")

		h.sendErrorResponse(c, err)

		return
	}

	log.Info("successful login")

	c.JSON(http.StatusOK, loginInfo)
}
