package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/capsa-gg/capsa/server/internal/domain/common"
	"github.com/capsa-gg/capsa/server/internal/server/bodies"
)

// UserSearch allows users search for resources
// @Summary 	Search for resources
// @Tags        UserSearch
// @Produce     json
// @Description	Allows users to search for resources that contain the search query as a substring
// @Security	JwtUser
// @Security	JwtAdmin
// @Param		search 	query		string 		true 	"SearchTerm"
// @Success		200		{object} 	bodies.SearchResults
// @Failure     400		{object}	bodies.ErrorResponse
// @Failure     404		{object}	bodies.ErrorResponse
// @Failure     500		{object}	bodies.ErrorResponse
// @Header		all		{string} 	X-Capsa-Server-Version		"Current Capsa Server version"
// @Router 		/user/search [get]
func (h Handlers) UserSearch(c *gin.Context) {
	log := h.logger.Named("UserSearch")

	search := c.Query("search")
	if search == "" {
		c.JSON(http.StatusBadRequest, bodies.ErrorResponse{Error: "search query is required"})

		return
	}

	log = log.With("search", search)
	log.Debug("attempting to search in database")

	res, err := common.PerformSearch(c, h.services, search)
	if err != nil {
		h.sendErrorResponse(c, err)

		return
	}

	log.Debug("search completed")

	c.JSON(http.StatusOK, res)
}
