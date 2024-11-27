package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/capsa-gg/capsa/server/constants"
	"github.com/capsa-gg/capsa/server/internal/domain/user"
	"github.com/capsa-gg/capsa/server/internal/domainerror"
	"github.com/capsa-gg/capsa/server/internal/server/bodies"
)

// ListAllUsers lists all users present in the database
// @Summary 	List all users
// @Tags        AdminUsers
// @Produce    	json
// @Description	Allows admins to fetch all users from the database
// @Security	JwtAdmin
// @Success		200		{array}		bodies.UserInfoResponse
// @Failure     400		{object}	bodies.ErrorResponse
// @Failure     404		{object}	bodies.ErrorResponse
// @Failure     500		{object}	bodies.ErrorResponse
// @Header		all		{string} 	X-Capsa-Server-Version			"Current Capsa Server version"
// @Router 		/admin/users [get]
func (h Handlers) ListAllUsers(c *gin.Context) {
	log := h.logger.Named("ListAllUsers")

	res, err := user.ListAllUsers(c, h.services)
	if err != nil {
		h.sendErrorResponse(c, err)

		return
	}

	log.Debugf("fetched %d items", len(res))

	c.JSON(http.StatusOK, res)
}

// CreateUser creates a new user
// @Summary 	Creates a new user
// @Tags        AdminUsers
// @Produce    	json
// @Description	Allows admins to create a new user
// @Security	JwtAdmin
// @Param		userinfo body 	bodies.UserCreateRequest 	true 	"UserCreateRequest"
// @Success		201		{object}	bodies.UserInfoResponse
// @Failure     400		{object}	bodies.ErrorResponse
// @Failure     404		{object}	bodies.ErrorResponse
// @Failure     500		{object}	bodies.ErrorResponse
// @Header		all		{string} 	X-Capsa-Server-Version			"Current Capsa Server version"
// @Router 		/admin/users [post]
func (h Handlers) CreateUser(c *gin.Context) {
	log := h.logger.Named("CreateUser")

	req, err := extractBodyJSON[bodies.UserCreateRequest](c, h.services)
	if err != nil {
		return // Error sent by extractBodyJSON
	}

	if req == nil {
		h.sendErrorResponse(c, jsonBodyExtractionError)

		return
	}

	log.Debug("body extracted")

	newUserRole, err := constants.UserRoleFromString(req.Role)
	if err != nil {
		h.sendErrorResponse(c, err)

		return
	}

	log = log.With("user_role", newUserRole)
	log.Debug("role extracted")

	newUserUUID, err := user.AddNewUser(c, h.services, req.Email, req.FirstName, req.LastName, newUserRole)
	if err != nil {
		h.sendErrorResponse(c, err)

		return
	}

	if newUserUUID == nil {
		h.sendErrorResponse(c, domainerror.New(domainerror.Unexpected, "no new user uuid", errors.New("newUserUUID is nil")))

		return
	}

	log = log.With("user_uuid", newUserRole)
	log.Debug("user added to database")

	newUser, err := user.GetUserByUUID(c, h.services, *newUserUUID)
	if err != nil {
		h.sendErrorResponse(c, err)

		return
	}

	c.JSON(http.StatusCreated, newUser)
}

// DeactivateUser deactivates a user.
// @Summary 	Deactivate a user's account
// @Tags        AdminUsers
// @Produce    	json
// @Description	Allows admins to deactivate a user
// @Security	JwtAdmin
// @Success		201		{object}	bodies.UserInfoResponse
// @Failure     400		{object}	bodies.ErrorResponse
// @Failure     404		{object}	bodies.ErrorResponse
// @Failure     500		{object}	bodies.ErrorResponse
// @Header		all		{string} 	X-Capsa-Server-Version			"Current Capsa Server version"
// @Router 		/admin/users/{useruuid}/activation [delete]
func (h Handlers) DeactivateUser(c *gin.Context) { //nolint:dupl // Unique enough for duplicate code
	log := h.logger.Named("DeactivateUser")

	userUUID, ok := getUserUUIDFromURI(c)
	if !ok {
		return // Response sent by getLogUUIDFromURI
	}

	_, userJwtUUID, ok := extractUserJwtClaimsFromContext(c, log)
	if !ok {
		return // Response sent by extractUserJwtClaimsFromContext
	}

	if userUUID == userJwtUUID {
		h.sendErrorResponse(c, domainerror.New(domainerror.NoPermission, "cannot change activation of your own user", errors.New("changing activation status of own user is not allowed")))

		return
	}

	log = log.With("user_uuid", userUUID)
	log.Debug("extracted user_uuid")

	err := user.DeactivateUser(c, h.services, userUUID)
	if err != nil {
		h.sendErrorResponse(c, err)

		return
	}

	log.Debug("user deactivated")

	userInfo, err := user.GetUserByUUID(c, h.services, userUUID)
	if err != nil {
		h.sendErrorResponse(c, err)

		return
	}

	c.JSON(http.StatusCreated, userInfo)
}

// ReactivateUser reactivates a user.
// @Summary 	Reactivate a deactivated user's account
// @Tags        AdminUsers
// @Produce    	json
// @Description	Allows admins to reactivate a deactivated user
// @Security	JwtAdmin
// @Success		201		{object}	bodies.UserInfoResponse
// @Failure     400		{object}	bodies.ErrorResponse
// @Failure     404		{object}	bodies.ErrorResponse
// @Failure     500		{object}	bodies.ErrorResponse
// @Header		all		{string} 	X-Capsa-Server-Version			"Current Capsa Server version"
// @Router 		/admin/users/{useruuid}/activation [post]
func (h Handlers) ReactivateUser(c *gin.Context) { //nolint:dupl // Unique enough for duplicate code
	log := h.logger.Named("DeactivateUser")

	userUUID, ok := getUserUUIDFromURI(c)
	if !ok {
		return // Response sent by getLogUUIDFromURI
	}

	_, userJwtUUID, ok := extractUserJwtClaimsFromContext(c, log)
	if !ok {
		return // Response sent by extractUserJwtClaimsFromContext
	}

	if userUUID == userJwtUUID {
		h.sendErrorResponse(c, domainerror.New(domainerror.NoPermission, "cannot change activation of your own user", errors.New("changing activation status of own user is not allowed")))

		return
	}

	log = log.With("user_uuid", userUUID)
	log.Debug("extracted user_uuid")

	err := user.ReactivateUser(c, h.services, userUUID)
	if err != nil {
		h.sendErrorResponse(c, err)

		return
	}

	log.Debug("user reactivated")

	userInfo, err := user.GetUserByUUID(c, h.services, userUUID)
	if err != nil {
		h.sendErrorResponse(c, err)

		return
	}

	c.JSON(http.StatusCreated, userInfo)
}

// UpdateUser updates a user.
// @Summary 	Update a user's info in the database
// @Tags        AdminUsers
// @Produce    	json
// @Description	Allows admins to update a user
// @Security	JwtAdmin
// @Param		userinfo body 	bodies.UserUpdateRequest 	true 	"UserUpdateRequest"
// @Success		201		{object}	bodies.UserInfoResponse
// @Failure     400		{object}	bodies.ErrorResponse
// @Failure     404		{object}	bodies.ErrorResponse
// @Failure     500		{object}	bodies.ErrorResponse
// @Header		all		{string} 	X-Capsa-Server-Version			"Current Capsa Server version"
// @Router 		/admin/users/{useruuid} [put]
func (h Handlers) UpdateUser(c *gin.Context) {
	log := h.logger.Named("DeactivateUser")

	userUUID, ok := getUserUUIDFromURI(c)
	if !ok {
		return // Response sent by getLogUUIDFromURI
	}

	log = log.With("user_uuid", userUUID)
	log.Debug("extracted user_uuid")

	_, userJwtUUID, ok := extractUserJwtClaimsFromContext(c, log)
	if !ok {
		return // Response sent by extractUserJwtClaimsFromContext
	}

	req, err := extractBodyJSON[bodies.UserUpdateRequest](c, h.services)
	if err != nil {
		return // Error sent by extractBodyJSON
	}

	if req == nil {
		h.sendErrorResponse(c, domainerror.New(domainerror.Unexpected, "nil arrgument", errors.New("userInfo is nil")))

		return
	}

	if userUUID == userJwtUUID {
		if req.Role != constants.UserRoleAdmin { // User must be an admin to reach this endpoint
			h.sendErrorResponse(c, domainerror.New(domainerror.InvalidArgument, "cannot change your own role", errors.New("changing role of own user is not allowed")))

			return
		}
	}

	log.Debug("validation succeeded, updating user")

	err = user.UpdateUser(c, h.services, userUUID, req)
	if err != nil {
		h.sendErrorResponse(c, err)

		return
	}

	log.Debug("user updated")

	userInfo, err := user.GetUserByUUID(c, h.services, userUUID)
	if err != nil {
		h.sendErrorResponse(c, err)

		return
	}

	c.JSON(http.StatusCreated, userInfo)
}
