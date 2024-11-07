package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/lucianonooijen/capsa/server/internal/domain/logs"
)

// LogsList allows users to fetch available logs from the database
// @Summary 	Log listing
// @Tags        UserLogs
// @Produce    	json
// @Description	Allows users to fetch available logs from the database
// @Security	JwtUser
// @Success		200		{array}		entities.LogOverview
// @Failure     400		{object}	bodies.ErrorResponse
// @Failure     404		{object}	bodies.ErrorResponse
// @Failure     500		{object}	bodies.ErrorResponse
// @Header		all		{string} 	X-Capsa-Server-Version			"Current Capsa Server version"
// @Router 		/user/logs [get]
func (h Handlers) LogsList(c *gin.Context) {
	log := h.logger.Named("LogsList")

	res, err := logs.GetAllLogsOverview(h.services)
	if err != nil {
		h.sendErrorResponse(c, err)

		return
	}

	log.Debugf("fetched %d items", len(res))

	c.JSON(http.StatusOK, res)
}
