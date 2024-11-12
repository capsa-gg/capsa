package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/capsa-gg/capsa/server/internal/server/bodies"
)

// Jwks does not have Swagger documentation available as it does not live under the /v1 endpoints.
func (h Handlers) Jwks(c *gin.Context) {
	log := h.logger.Named("Jwks")

	pubKey, err := h.services.Token.GetPublicKey()
	if err != nil {
		log.Errorf("errorr getting public key information: %s", err)

		c.JSON(http.StatusInternalServerError, bodies.ErrorResponse{Error: "cannot get public key information"})

		return
	}

	c.Header("content-type", "application/json")

	// Manually writing JSON for compatibility with the []byte from pubKey
	c.Writer.WriteString(`{"keys":[`) //nolint:errcheck // This is fine
	c.Writer.Write(pubKey)            //nolint:errcheck // This is fine
	c.Writer.WriteString(`]}`)        //nolint:errcheck // This is fine

	c.Status(http.StatusOK)
}
