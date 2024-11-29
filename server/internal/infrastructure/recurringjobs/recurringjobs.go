package recurringjobs

import (
	"context"
	"fmt"

	"github.com/robfig/cron/v3"
	"go.uber.org/zap"

	"github.com/capsa-gg/capsa/server/internal/domain/logs"
	"github.com/capsa-gg/capsa/server/internal/entities"
	"github.com/capsa-gg/capsa/server/internal/interactor"
)

// RecurringJobs performs recurring jobs on a schedule.
type RecurringJobs struct {
	jobs     *cron.Cron
	logger   *zap.SugaredLogger
	services *interactor.Services
}

// New initializes a new RecurringJobs with the required jobs added.
func New(c *entities.Config, s *interactor.Services) (*RecurringJobs, error) {
	logger := c.RootLogger.Named("RecurringJobs").Sugar()
	cj := cron.New()

	// Every midnight, run cleaning job for log removal
	_, err := cj.AddFunc("0 0 * * *", func() {
		log := logger.Named("DailyLogCleanFunc")
		log.Info("Starting daily log cleaning")

		err := logs.DeleteLogsOverRetention(context.Background(), s)
		if err != nil {
			log.Errorf("error from log deletion: %s", err)
		}

		log.Info("Log cleaning finished")
	})
	if err != nil {
		return nil, fmt.Errorf("error adding function for daily log cleaning: %w", err)
	}

	recurringJobs := RecurringJobs{logger: logger, services: s, jobs: cj}

	return &recurringJobs, nil
}

// Start starts running the scheduled jobs.
func (rc *RecurringJobs) Start() {
	rc.jobs.Start()
}

// Stop stops running the scheduled jobs.
func (rc *RecurringJobs) Stop() {
	rc.jobs.Stop()
}
