package logs

import (
	"context"
	"encoding/json"

	"github.com/jackc/pgx/v5/pgtype"

	"github.com/capsa-gg/capsa/server/constants"
	"github.com/capsa-gg/capsa/server/internal/entities"
	"github.com/capsa-gg/capsa/server/internal/interactor"
	"github.com/capsa-gg/capsa/server/internal/util"
)

// GetAllLogsOverview fetches the high-level overview of all available logs.
// TODO: pagination.
func GetAllLogsOverview(ctx context.Context, s *interactor.Services) ([]entities.LogOverview, error) {
	log := s.GetDomainLogger("logs", "GetAllLogsOverview")

	// Get from database
	rows, err := s.Database.ListAvailableLogs(ctx, pgtype.UUID{Valid: false})
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

		if err = json.Unmarshal(rows[i].CategoriesCount, &logsAvailable[i].CategoriesCounts); err != nil {
			logLoop.Errorf("cannot convert CategoriesCount to map[string]int: %s", err)
		}

		if err = json.Unmarshal(rows[i].SeveritiesCount, &logsAvailable[i].SeveritiesCounts); err != nil {
			logLoop.Errorf("cannot convert SeveritiesCounts to map[string]int: %s", err)
		}
	}

	return logsAvailable, nil
}
