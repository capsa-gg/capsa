package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/capsa-gg/capsa/server/internal/domain/admin"
)

// EnvironmentsList allows users to fetch all available environments from the database
// @Summary 	Environment listing
// @Tags        UserEnvironments
// @Produce    	json
// @Description	Allows users to fetch all available environments from the database
// @Security	JwtUser
// @Success		200		{array}		entities.TitleEnvironment
// @Failure     400		{object}	bodies.ErrorResponse
// @Failure     404		{object}	bodies.ErrorResponse
// @Failure     500		{object}	bodies.ErrorResponse
// @Header		all		{string} 	X-Capsa-Server-Version			"Current Capsa Server version"
// @Router 		/user/environments [get]
func (h Handlers) EnvironmentsList(c *gin.Context) {
	log := h.logger.Named("LogsList")

	res, err := admin.ListAllTitlesAndEnvironments(c, h.services)
	if err != nil {
		h.sendErrorResponse(c, err)

		return
	}

	log.Debugf("fetched %d items", len(res))

	c.JSON(http.StatusOK, res)
}
