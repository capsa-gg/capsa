package logs

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"

	"github.com/capsa-gg/capsa/server/internal/interactor"
)

// DeleteLogsOverRetention uses the maximum log lifetime defined in the config to generate a list of all logs to be deleted, and then deletes them.
func DeleteLogsOverRetention(ctx context.Context, s *interactor.Services) error {
	now := time.Now()
	log := s.GetDomainLogger("logs", "DeleteLogsOverRetention").With("start", now)

	if s.Config.LogRetentionDays < 1 {
		log.Warnf("log retention days in config is set to %d, which is <1, not removing logs", s.Config.LogRetentionDays)

		return nil
	}

	deleteBefore := now.AddDate(0, 0, -s.Config.LogRetentionDays)

	// Just a sanity check to prevent all logs from being wiped
	if deleteBefore.After(now) {
		log.Error("delete before is after now!")

		return errors.New("delete before is after now")
	}

	// Get all logs that are before the deleteBefore timestamp
	logIDs, err := s.Database.GetLogsBeforeTimestamp(ctx, deleteBefore)
	if err != nil {
		return errors.New("error fetching logs from before timestamp")
	}

	// Exit early if we have no logs
	if len(logIDs) == 0 {
		log.Warn("no logs for deletion found")

		return nil
	}

	log = log.With("count_logids", len(logIDs))
	log.Info("start removing logs")

	failures := 0

	for _, logID := range logIDs {
		logLoop := log.Named("deleteLogsAndChunks").Named(strconv.FormatInt(logID, 10))
		start := time.Now()

		err = deleteLogsAndChunks(ctx, s, logLoop, logID)
		if err != nil {
			failures++

			logLoop.Errorf("error deleting log and chunks: %s", err)
		} else {
			logLoop.Info("log successfully removed")
		}

		end := time.Now()
		took := end.Sub(start).String()

		log.Debugf("took %s to process", took)
	}

	end := time.Now()
	took := end.Sub(now).String()

	log = log.With("count_failures", failures)
	log.Infof("log removal concluded after %s", took)

	return nil
}

// deleteLogsAndChunks attempts to delete logs from the database and chunks from the blobstorage.
// This is done on a best-effort basis for items in blob storage, only database errors are returned as error.
func deleteLogsAndChunks(ctx context.Context, s *interactor.Services, log *zap.SugaredLogger, logID int64) error {
	tx, err := s.DBPool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("error starting database transaction: %w", err)
	}
	defer func(tx pgx.Tx, ctx context.Context) {
		errRb := tx.Rollback(ctx)
		if errRb != nil {
			log.Warnf("error rolling back transaction: %s", errRb)
		}
	}(tx, ctx)

	qtx := s.Database.WithTx(tx)

	chunks, err := qtx.GetLogChunksForLog(ctx, logID)
	if err != nil {
		return fmt.Errorf("cannot load chunks for log: %w", err)
	}

	// Delete from blob storage
	if len(chunks) == 0 {
		log.Warn("no chunks for log, not removing from blob storage")
	} else {
		for i, c := range chunks {
			path := c.BlobPath
			logLoop := log.
				With("chunk", fmt.Sprintf("%d/%d", i+1, len(chunks))).
				With("path", path)

			err = s.LogChunks.DeleteChunk(path)
			if err != nil {
				logLoop.Errorf("error deleting chunk: %s", err)
			} else {
				logLoop.Info("chunk deleted")
			}
		}
	}

	// Delete from database
	err = qtx.DeleteLogAndLinkedResources(ctx, logID)
	if err != nil {
		return fmt.Errorf("error deleting log in database: %w", err)
	}

	// Commit
	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("error committing changes to database: %w", err)
	}

	return nil
}
