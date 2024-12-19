package logchunk

import (
	"context"
	"errors"
	"slices"
	"strings"

	"go.uber.org/zap"

	"github.com/capsa-gg/capsa/server/internal/data/database"
	"github.com/capsa-gg/capsa/server/internal/domainerror"
	"github.com/capsa-gg/capsa/server/internal/entities"
	"github.com/capsa-gg/capsa/server/internal/interactor"
)

type filteredLineLoader struct {
	log               *zap.SugaredLogger
	services          *interactor.Services
	filters           LogStreamLineFilters
	chunks            []database.LogsChunk
	lineCounter       int
	nextLineIndex     int
	currentChunkIndex int
	currentChunkLines []string
}

func newFilteredLineLoader(log *zap.SugaredLogger, s *interactor.Services, filters LogStreamLineFilters, chunks []database.LogsChunk) entities.FilteredLineLoader {
	return &filteredLineLoader{
		log:               log,
		services:          s,
		filters:           filters,
		chunks:            chunks,
		lineCounter:       0,
		currentChunkIndex: 0,
		currentChunkLines: []string{},
		nextLineIndex:     0,
	}
}

// GenerateFilteredLineLoaderForLog fetches the metadata of log chunks for a given logID and uses this to generate an entities.FilteredLineLoader.
// This function is to be used for merged logs. The functionality is close to StreamFilteredLogChunks.
// NOTE: the ID is not validated, this should be done in the calling function.
func GenerateFilteredLineLoaderForLog(ctx context.Context, s *interactor.Services, logID int64, filters LogStreamLineFilters) (entities.FilteredLineLoader, error) {
	log := s.GetDomainLogger("logs", "StreamLogChunks").With("log_id", logID)

	log.Debug("start chunk streaming")

	// Get chunks from database
	chunks, err := s.Database.GetLogChunksForLog(ctx, logID)
	if err != nil {
		return nil, domainerror.NewFromDatabaseError(err)
	}

	loader := newFilteredLineLoader(log.Named("filteredLineLoader"), s, filters, chunks)

	return loader, nil
}

func (fll *filteredLineLoader) HasNextLine() (bool, error) {
	log := fll.log.Named("HasNextLine").With("i_line_in_chunk", fll.nextLineIndex, "len_chunk_lines", len(fll.currentChunkLines))

	if fll.nextLineIndex+1 > len(fll.currentChunkLines) {
		log.Debug("no more data in chunk, loading next chunk")

		fetchedNext, err := fll.loadNextBlob()
		if err != nil {
			return false, err
		}

		// Last blob
		if !fetchedNext {
			return false, nil
		}

		// Recursively call itself to check for next line
		return fll.HasNextLine()
	}

	return true, nil
}

func (fll *filteredLineLoader) ReadNextLineMetadata() (lineMetadata *entities.LogChunkLineMetadata, err error) {
	log := fll.log.Named("ReadNextLineMetadata")

	hasNext, err := fll.HasNextLine()

	if err != nil {
		return nil, err
	}

	if !hasNext {
		log.Info("No next line present")

		return nil, nil //nolint:nilnil // this is well documented in the interface.
	}

	lineContents := fll.currentChunkLines[fll.nextLineIndex]

	if lineContents == "" {
		return nil, domainerror.New(domainerror.Unexpected, "line contents empty", errors.New("line contents empty"))
	}

	// Remove absolute line prefix for metadata extraction, slightly hacky maybe
	indexCloseBracketAbsLineNum := strings.IndexRune(lineContents, '}')
	if indexCloseBracketAbsLineNum > -1 {
		lineContents = lineContents[indexCloseBracketAbsLineNum+1:]
	}

	metadata, err := ExtractMetadataFromLine([]byte(lineContents))
	if err != nil {
		log.Errorf("error extracting line metadata for line '%s': %s", lineContents, err)
	}

	return &metadata, err
}

func (fll *filteredLineLoader) GetNextLine() (logLine *string, err error) {
	log := fll.log.Named("GetNextLine")

	hasNext, err := fll.HasNextLine()

	if err != nil {
		return nil, err
	}

	if !hasNext {
		log.Info("No next line present")

		return nil, nil //nolint:nilnil // this is well documented in the interface.
	}

	lineContents := fll.currentChunkLines[fll.nextLineIndex]

	if lineContents == "" {
		return nil, nil //nolint:nilnil // this is well documented in the interface.
	}

	fll.nextLineIndex++

	return &lineContents, nil
}

// Boolean indicates whether the next chunk has been loaded.
func (fll *filteredLineLoader) loadNextBlob() (bool, error) {
	log := fll.log.Named("loadNextBlob")

	for {
		logLoop := log.With("i_current_chunk", fll.currentChunkIndex)

		// Check if we have any more chunks
		if fll.currentChunkIndex+1 > len(fll.chunks) {
			logLoop.Debug("no more chunks")

			return false, nil
		}

		chunk := fll.chunks[fll.currentChunkIndex]
		fll.currentChunkIndex++ // Increment chunk

		shouldStreamChunk := fll.filters.shouldStreamChunk(logChunkMetadata{
			// Note: we omit some fields here, as they are not used
			SeveritiesCount: chunk.SeverityCounts,
			CategoriesCount: chunk.CategoryCounts,
		})

		if !shouldStreamChunk {
			fll.lineCounter += int(chunk.LineCount) // Needed for absolute line numbers

			logLoop.Debug("shouldStreamChunk returned false, skipping chunk processing/streaming")

			continue
		}

		chunkText, err := fll.services.LogChunks.GetChunk(chunk.BlobPath)
		if err != nil {
			return false, domainerror.New(domainerror.Unexpected, "error getting chunk from storage", err)
		}

		filteredLines := filterLinesForChunk(chunkText, fll.filters, &fll.lineCounter)

		if len(filteredLines) == 0 {
			logLoop.Warn("no contents after filtering chunk, not streaming data")

			continue
		}

		filteredLinesSlice := strings.Split(string(filteredLines), "\n")
		filteredNonEmptyLine := slices.DeleteFunc(filteredLinesSlice, func(s string) bool { return s == "" })

		fll.currentChunkLines = filteredNonEmptyLine
		fll.nextLineIndex = 0

		logLoop.Debug("fetched chunk contents")

		break
	}

	return true, nil
}
