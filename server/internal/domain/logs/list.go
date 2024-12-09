package logs

import (
	"context"
	"encoding/json"

	"github.com/capsa-gg/capsa/server/constants"
	"github.com/capsa-gg/capsa/server/internal/data/database"
	"github.com/capsa-gg/capsa/server/internal/domainerror"
	"github.com/capsa-gg/capsa/server/internal/interactor"
	"github.com/capsa-gg/capsa/server/internal/server/bodies"
	"github.com/capsa-gg/capsa/server/internal/util"
)

const getLogsLimitCount = 1000

// GetLogs fetches the high-level overview of all available logs.
func GetLogs(ctx context.Context, s *interactor.Services, filters ListFilters) ([]bodies.LogInfo, bool, error) { //nolint:gocyclo // This is more readable than breaking it up
	log := s.GetDomainLogger("logs", "GetLogs").With("filters", filters)

	params := database.ListAvailableLogsParams{
		FilterByLogUuid:     nil,
		FilterByEnvironment: filters.Environment,
		FilterByLogtype:     database.NullLogClientType{Valid: false},
		FilterByPlatform:    filters.Platform,
		Fetchlimit:          getLogsLimitCount + 1, // Using +1 here to check for hasMore
	}

	if filters.LogType != nil {
		params.FilterByLogtype = database.NullLogClientType{
			LogClientType: database.LogClientType(*filters.LogType), // Safe conversion
			Valid:         true,
		}
	}

	// Get from database
	rows, err := s.Database.ListAvailableLogs(ctx, params)
	if err != nil {
		return nil, false, domainerror.NewFromDatabaseError(err)
	}

	logsAvailable := make([]bodies.LogInfo, 0, len(rows))

	for i := range rows {
		if i >= getLogsLimitCount {
			break
		}

		logLoop := log.With("log_uuid", rows[i].LogUuid)

		info := bodies.LogInfo{
			LogUUID:        rows[i].LogUuid,
			LogType:        constants.LogType(rows[i].LogType), // safe conversion
			Title:          rows[i].Title,
			Environment:    rows[i].Environment,
			Platform:       rows[i].Platform,
			LineCount:      rows[i].LineCount,
			ChunkCount:     rows[i].ChunkCount,
			LinkedLogCount: rows[i].LinkCount,
		}

		earliest := rows[i].Earliest

		if earliest != nil {
			earliestTS, err := util.ExtractTimeFromAny(earliest)
			if err != nil {
				logLoop.Errorf("could not convert earlist value %#v to time: %s", earliest, err)
			} else {
				info.TimestampFirstLine = &earliestTS
			}
		}

		last := rows[i].Last

		if last != nil {
			lastTS, err := util.ExtractTimeFromAny(last)
			if err != nil {
				logLoop.Errorf("could not convert last value %#v to time: %s", last, err)
			} else {
				info.TimestampLastLine = &lastTS
			}
		}

		if err = json.Unmarshal(rows[i].CategoriesCount, &info.CategoriesCounts); err != nil {
			logLoop.Errorf("cannot convert CategoriesCount to map[string]int: %s", err)
		}

		if err = json.Unmarshal(rows[i].SeveritiesCount, &info.SeveritiesCounts); err != nil {
			logLoop.Errorf("cannot convert SeveritiesCounts to map[string]int: %s", err)
		}

		logsAvailable = append(logsAvailable, info)
	}

	hasMore := len(rows) > getLogsLimitCount

	return logsAvailable, hasMore, nil
}
