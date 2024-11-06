package admin

import (
	"context"
	"fmt"

	"github.com/lucianonooijen/capsa/server/internal/interactor"
)

// AddNewTitle adds a new title to the database.
func AddNewTitle(s *interactor.Services, title string) error {
	log := s.GetDomainLogger("admin", "AddNewTitle").With("title", title)
	ctx := context.TODO()

	log.Debug("attempting to add title")

	err := s.Database.AddTitle(ctx, title)
	if err != nil {
		return fmt.Errorf("error creating title: %w", err)
	}

	log.Info("title added")

	return nil
}
