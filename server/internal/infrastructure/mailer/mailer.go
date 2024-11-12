package mailer

import (
	"go.uber.org/zap"

	"github.com/capsa-gg/capsa/server/internal/entities"
)

// Mailer is the generic package used for sending transactional emails.
// The actual emails are generated in the data/emails package.
type Mailer struct {
	logger *zap.Logger
	config *entities.Config
}

// New generates a new Mailer instance.
func New(c *entities.Config) entities.EmailSender {
	log := c.RootLogger.Named("mailer")

	return &Mailer{
		logger: log,
		config: c,
	}
}
