package admin

import (
	"context"

	"github.com/lucianonooijen/capsa/server/internal/data/database"
	"github.com/lucianonooijen/capsa/server/internal/entities"
	"github.com/lucianonooijen/capsa/server/internal/interactor"
)

// AddNewEnvironment adds a new environment to the database for a given title.
func AddNewEnvironment(s *interactor.Services, title, env string) error {
	log := s.GetDomainLogger("admin", "AddNewEnvironment").With("title", title, "environment", env)
	ctx := context.TODO()

	log.Debug("attempting to add environment")

	titleInfo, err := s.Database.GetTitleByName(ctx, title)
	if err != nil {
		log.Warnf("cannot get title: %s", err)

		return entities.NewDomainErrorFromDatabaseError(err)
	}

	log = log.With("title_id", titleInfo.ID)
	log.Debug("fetched title data")

	err = s.Database.AddEnvironment(ctx, database.AddEnvironmentParams{
		Title: titleInfo.ID,
		Name:  env,
	})

	if err != nil {
		log.Warnf("cannot add environment: %s", err)

		return entities.NewDomainErrorFromDatabaseError(err)
	}

	log.Info("environment added")

	return nil
}

// ListAllTitlesAndEnvironments lists all titles and environments, with the associated keys.
func ListAllTitlesAndEnvironments(s *interactor.Services) ([]database.GetAllEnvironmentsAndTitlesRow, error) {
	log := s.GetDomainLogger("admin", "ListAllTitlesAndEnvironments")
	ctx := context.TODO()

	res, err := s.Database.GetAllEnvironmentsAndTitles(ctx)
	if err != nil {
		log.Warnf("cannot get titles and environments: %s", err)

		return nil, entities.NewDomainErrorFromDatabaseError(err)
	}

	log.Infof("fetched %d items", len(res))

	return res, err
}
