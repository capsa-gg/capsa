package admin

import (
	"context"

	"github.com/capsa-gg/capsa/server/internal/domainerror"
	"github.com/capsa-gg/capsa/server/internal/interactor"
	"github.com/capsa-gg/capsa/server/internal/server/bodies"
)

// ListAllTitles returns all titles in the database.
func ListAllTitles(ctx context.Context, s *interactor.Services) ([]bodies.TitleResponse, error) {
	log := s.GetDomainLogger("admin", "ListAllTitles")

	log.Debug("fetching all titles")

	titlesDB, err := s.Database.ListTitles(ctx)
	if err != nil {
		return nil, domainerror.NewFromDatabaseError(err)
	}

	titles := make([]bodies.TitleResponse, 0, len(titlesDB))

	for i := range titlesDB {
		titleDB := titlesDB[i]

		title := bodies.TitleResponse{
			Title:     titleDB.Name,
			CreatedOn: titleDB.CreatedOn,
		}

		titles = append(titles, title)
	}

	return titles, nil
}

// AddNewTitle adds a new title to the database.
func AddNewTitle(ctx context.Context, s *interactor.Services, title string) error {
	log := s.GetDomainLogger("admin", "AddNewTitle").With("title", title)

	log.Debug("attempting to add title")

	err := s.Database.AddTitle(ctx, title)
	if err != nil {
		log.Warnf("cannot add title: %s", err)

		return domainerror.NewFromDatabaseError(err)
	}

	log.Info("title added")

	return nil
}
