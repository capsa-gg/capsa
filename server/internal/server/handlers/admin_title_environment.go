package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/capsa-gg/capsa/server/internal/domain/admin"
	"github.com/capsa-gg/capsa/server/internal/domainerror"
	"github.com/capsa-gg/capsa/server/internal/server/bodies"
	"github.com/capsa-gg/capsa/server/internal/util"
)

func errorInvalidTitle(name string) error {
	return domainerror.New(domainerror.InvalidArgument, name+" is not a valid title name, must conform to a-zA-Z0-9", errors.New("title name invalid"))
}

func errorInvalidEnvironment(name string) error {
	return domainerror.New(domainerror.InvalidArgument, name+" is not a valid environment name, must conform to a-zA-Z0-9", errors.New("environment name invalid"))
}

// TitlesList allows admins to fetch titles from the database
// @Summary 	List all titles
// @Tags        AdminTitleEnvironment
// @Produce    	json
// @Description	Allows admins to add list all titles in the database
// @Security	JwtAdmin
// @Success		200		{array}		bodies.TitleResponse
// @Failure     400		{object}	bodies.ErrorResponse
// @Failure     404		{object}	bodies.ErrorResponse
// @Failure     500		{object}	bodies.ErrorResponse
// @Header		all		{string} 	X-Capsa-Server-Version			"Current Capsa Server version"
// @Router 		/admin/titles [get]
func (h Handlers) TitlesList(c *gin.Context) {
	log := h.logger.Named("TitlesList")

	titles, err := admin.ListAllTitles(c, h.services)
	if err != nil {
		h.sendErrorResponse(c, err)

		return
	}

	log.Debugf("fetched %d items", len(titles))

	c.JSON(http.StatusOK, titles)
}

// TitleAdd allows admins add a new title to the database
// @Summary 	Add new title
// @Tags        AdminTitleEnvironment
// @Produce    	json
// @Description	Allows admins to add new titles
// @Security	JwtAdmin
// @Param		title	body 		bodies.AddTitleRequest 	true 	"TitleAddRequest"
// @Success		201		{array}		bodies.TitleResponse
// @Failure     400		{object}	bodies.ErrorResponse
// @Failure     404		{object}	bodies.ErrorResponse
// @Failure     500		{object}	bodies.ErrorResponse
// @Header		all		{string} 	X-Capsa-Server-Version			"Current Capsa Server version"
// @Router 		/admin/titles [post]
func (h Handlers) TitleAdd(c *gin.Context) {
	log := h.logger.Named("TitleAdd")

	req, err := extractBodyJSON[bodies.AddTitleRequest](c, h.services)
	if err != nil {
		return // Error sent by extractBodyJSON
	}

	if req == nil {
		h.sendErrorResponse(c, domainerror.New(domainerror.Unexpected, "request body nil", errors.New("req is nil")))

		return
	}

	log = log.With("title_name", req.TitleName)
	log.Debug("extracted title name")

	// Validate name
	if !util.IsValidTitleOrEnvironmentName(req.TitleName) {
		h.sendErrorResponse(c, errorInvalidTitle(req.TitleName))

		return
	}

	log.Debug("title name is valid, adding title")

	// Add title
	err = admin.AddNewTitle(c, h.services, req.TitleName)
	if err != nil {
		h.sendErrorResponse(c, err)

		return
	}

	// Return all titles
	titles, err := admin.ListAllTitles(c, h.services)
	if err != nil {
		h.sendErrorResponse(c, err)

		return
	}

	log.Debugf("fetched %d items", len(titles))

	c.JSON(http.StatusCreated, titles)
}

// EnvironmentAdd allows admins to add an environment for an existing title
// @Summary 	Add new environment for an existing title
// @Tags        AdminTitleEnvironment
// @Produce    	json
// @Description	Allows admins to add new environments for an existing title
// @Security	JwtAdmin
// @Param		environment	body	bodies.AddEnvironmentRequest 	true 	"EnvironmentAddRequest"
// @Success		201		{array}		bodies.TitleEnvironmentResponse
// @Failure     400		{object}	bodies.ErrorResponse
// @Failure     404		{object}	bodies.ErrorResponse
// @Failure     500		{object}	bodies.ErrorResponse
// @Header		all		{string} 	X-Capsa-Server-Version			"Current Capsa Server version"
// @Router 		/admin/environments [post]
func (h Handlers) EnvironmentAdd(c *gin.Context) {
	log := h.logger.Named("TitleAdd")

	req, err := extractBodyJSON[bodies.AddEnvironmentRequest](c, h.services)
	if err != nil {
		return // Error sent by extractBodyJSON
	}

	if req == nil {
		h.sendErrorResponse(c, domainerror.New(domainerror.Unexpected, "request body nil", errors.New("req is nil")))

		return
	}

	log = log.With("title_name", req.TitleName).With("environment_name", req.EnvironmentName)
	log.Debug("extracted title and environment names")

	// Validate names
	if !util.IsValidTitleOrEnvironmentName(req.TitleName) {
		h.sendErrorResponse(c, errorInvalidTitle(req.TitleName))

		return
	}

	if !util.IsValidTitleOrEnvironmentName(req.EnvironmentName) {
		h.sendErrorResponse(c, errorInvalidEnvironment(req.EnvironmentName))

		return
	}

	log.Debug("title and environment names valid, adding environment")

	// Add environment
	err = admin.AddNewEnvironment(c, h.services, req.TitleName, req.EnvironmentName)
	if err != nil {
		h.sendErrorResponse(c, err)

		return
	}

	// Return all titles and environments
	res, err := admin.ListAllTitlesAndEnvironments(c, h.services)
	if err != nil {
		h.sendErrorResponse(c, err)

		return
	}

	log.Debugf("fetched %d items", len(res))

	c.JSON(http.StatusCreated, res)
}
