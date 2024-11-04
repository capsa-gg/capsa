package admin

import (
	"context"
	"fmt"

	"github.com/lucianonooijen/capsa/server/internal/data/database"
	"github.com/lucianonooijen/capsa/server/internal/interactor"
)

// AddNewTitle adds a new title to the database.
func AddNewTitle(s *interactor.Services, title string) error {
	log := s.GetDomainLogger("admin", "AddNewTitle").With("title", title)
	ctx := context.TODO()

	log.Debugf("attempting to add title")

	err := s.Database.AddTitle(ctx, title)
	if err != nil {
		return fmt.Errorf("error creating title: %w", err)
	}

	log.Infof("title added")

	return nil
}

// AddNewEnvironment adds a new environment to the database for a given title.
func AddNewEnvironment(s *interactor.Services, title, env string) error {
	log := s.GetDomainLogger("admin", "AddNewEnvironment").With("title", title, "environment", env)
	ctx := context.TODO()

	log.Debugf("attempting to add environment")

	titleInfo, err := s.Database.GetTitleByName(ctx, title)
	if err != nil {
		return fmt.Errorf("error getting title %s by name: %w", title, err)
	}

	log = log.With("title_id", titleInfo.ID)
	log.Debugf("fetched title data")

	err = s.Database.AddEnvironment(ctx, database.AddEnvironmentParams{
		Title: titleInfo.ID,
		Name:  env,
	})

	if err != nil {
		return fmt.Errorf("error creating environment: %w", err)
	}

	log.Infof("environment added")

	return nil
}
