package emails

import (
	"fmt"
	"time"

	"github.com/matcornic/hermes/v2"
	"go.uber.org/zap"

	"github.com/capsa-gg/capsa/server/constants"
	"github.com/capsa-gg/capsa/server/internal/entities"
)

// Emails is used to send transactional emails using an entities.Mailer instance.
type Emails struct {
	logger *zap.SugaredLogger
	config *entities.Config
	mailer entities.EmailSender
	hermes *hermes.Hermes
}

// New returns a new instance of Emails.
func New(c *entities.Config, mailer entities.EmailSender) *Emails {
	logger := c.RootLogger.Named("Emails").Sugar()

	staticFileURLBase := fmt.Sprintf("https://%s/v1%s", c.ServerHostname, constants.APIStaticPath)

	hermesInstance := hermes.Hermes{
		Product: hermes.Product{
			Name:        "Capsa Server",
			Link:        fmt.Sprintf("%s://%s", c.ServerProtocol, c.WebappHostname), // Note: this assumes that the webapp is running on the same protocol as the server
			Logo:        fmt.Sprintf("%s/logo-with-by.png", staticFileURLBase),      //nolint:perfsprint // More readable this way
			Copyright:   fmt.Sprintf("Copyright © %d Capsa. All rights reserved.", time.Now().Year()),
			TroubleText: "If the button '{ACTION}' does not work, copy and paste the URL below.",
		},
	}

	return &Emails{
		logger: logger,
		config: c,
		mailer: mailer,
		hermes: &hermesInstance,
	}
}

func (e Emails) generatePasswordResetLink(token string) string {
	// Note: this assumes that the webapp is running on the same protocol as the server
	return fmt.Sprintf("%s://%s/auth/password-reset?resetToken=%s", e.config.ServerProtocol, e.config.WebappHostname, token)
}
