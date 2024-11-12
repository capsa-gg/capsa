package admin

import (
	"context"

	"github.com/capsa-gg/capsa/server/internal/entities"

	"github.com/capsa-gg/capsa/server/internal/interactor"
)

// AddNewTitle adds a new title to the database.
func AddNewTitle(s *interactor.Services, title string) error {
	log := s.GetDomainLogger("admin", "AddNewTitle").With("title", title)
	ctx := context.TODO()

	log.Debug("attempting to add title")

	err := s.Database.AddTitle(ctx, title)
	if err != nil {
		log.Warnf("cannot add title: %s", err)

		return entities.NewDomainErrorFromDatabaseError(err)
	}

	log.Info("title added")

	return nil
}
