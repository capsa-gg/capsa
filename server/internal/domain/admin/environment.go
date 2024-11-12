package admin

import (
	"context"

	"github.com/capsa-gg/capsa/server/internal/data/database"
	"github.com/capsa-gg/capsa/server/internal/entities"
	"github.com/capsa-gg/capsa/server/internal/interactor"
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
func ListAllTitlesAndEnvironments(s *interactor.Services) ([]entities.TitleEnvironment, error) {
	log := s.GetDomainLogger("admin", "ListAllTitlesAndEnvironments")
	ctx := context.TODO()

	res, err := s.Database.ListAllEnvironmentsAndTitles(ctx)
	if err != nil {
		log.Warnf("cannot get titles and environments: %s", err)

		return nil, entities.NewDomainErrorFromDatabaseError(err)
	}

	log.Infof("fetched %d items", len(res))

	titleEnvironments := make([]entities.TitleEnvironment, len(res))
	for i, te := range res {
		titleEnvironments[i].Title = te.Title
		titleEnvironments[i].EnvironmentName = te.Environment
		titleEnvironments[i].EnvironmentKey = te.EnvironmentKey
	}

	return titleEnvironments, err
}
