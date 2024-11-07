package logs

import (
	"context"

	"github.com/lucianonooijen/capsa/server/constants"
	"github.com/lucianonooijen/capsa/server/internal/entities"
	"github.com/lucianonooijen/capsa/server/internal/interactor"
	"github.com/lucianonooijen/capsa/server/internal/util"
)

// GetAllLogsOverview fetches the high-level overview of all available logs.
// TODO: pagination.
func GetAllLogsOverview(s *interactor.Services) ([]entities.LogOverview, error) {
	log := s.GetDomainLogger("logs", "GetAllLogsOverview")
	ctx := context.TODO()

	// Get from database
	rows, err := s.Database.ListAllAvailableLogs(ctx)
	if err != nil {
		return nil, entities.NewDomainErrorFromDatabaseError(err)
	}

	logsAvailable := make([]entities.LogOverview, len(rows))

	for i := range rows {
		logLoop := log.With("log_uuid", rows[i].LogUuid)

		logsAvailable[i].LogUUID = rows[i].LogUuid
		logsAvailable[i].LogType = constants.LogType(rows[i].LogType) // safe conversion
		logsAvailable[i].Platform = rows[i].Platform
		logsAvailable[i].LineCount = rows[i].LineCount

		earliest := rows[i].Earliest

		if earliest != nil {
			earliestTS, err := util.ExtractTimeFromAny(earliest)
			if err != nil {
				logLoop.Errorf("could not convert earlist value %#v to time: %s", earliest, err)
			} else {
				logsAvailable[i].TimestampFirstLine = &earliestTS
			}
		}

		last := rows[i].Last

		if last != nil {
			lastTS, err := util.ExtractTimeFromAny(last)
			if err != nil {
				logLoop.Errorf("could not convert last value %#v to time: %s", last, err)
			} else {
				logsAvailable[i].TimestampLastLine = &lastTS
			}
		}
	}

	return logsAvailable, nil
}
