package emails

import (
	"fmt"
	"time"

	"github.com/matcornic/hermes/v2"
	"go.uber.org/zap"

	"github.com/lucianonooijen/capsa/server/internal/entities"
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

	hermesInstance := hermes.Hermes{
		Product: hermes.Product{
			Name: "Capsa Server",
			Link: "https://capsa.lucianonooijen.com",
			//Logo:        fmt.Sprintf("%s/logo.png", staticFileURLBase), // TODO: Serve static files
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
	// TODO: read http/https from config
	return fmt.Sprintf("https://%s/auth/password-reset?resetToken=%s", e.config.WebappHostname, token)
}
