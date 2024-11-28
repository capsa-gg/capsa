package common

import (
	"context"

	"github.com/capsa-gg/capsa/server/internal/data/database"
	"github.com/capsa-gg/capsa/server/internal/domainerror"
	"github.com/capsa-gg/capsa/server/internal/interactor"
	"github.com/capsa-gg/capsa/server/internal/server/bodies"
)

const searchLimit = 1000

// PerformSearch performs a database search.
func PerformSearch(ctx context.Context, s *interactor.Services, searchTerm string) (*bodies.SearchResults, error) {
	log := s.GetDomainLogger("common", "PerformSearch").With("search_term", searchTerm)

	log.Debug("start search")

	matches, err := s.Database.SearchResources(ctx, database.SearchResourcesParams{
		Search: searchTerm,
		Limit:  searchLimit + 1, // Increment by 1, so we can check for HasMore
	})
	if err != nil {
		return nil, domainerror.NewFromDatabaseError(err)
	}

	log = log.With("len_matches", len(matches))
	log.Debug("received results")

	result := bodies.SearchResults{
		HasMore: len(matches) > searchLimit,
		Results: make([]bodies.SearchResult, 0, len(matches)),
	}

	for i := range matches {
		if i >= searchLimit {
			break
		}

		match := matches[i]
		result.Results = append(result.Results, bodies.SearchResult{
			Type:        match.TableName,
			Match:       match.Identifier,
			Description: match.Description,
			Details:     match.Details,
		})
	}

	return &result, nil
}
