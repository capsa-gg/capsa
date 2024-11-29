package logchunks

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/capsa-gg/capsa/server/internal/entities"
)

// LogChunks is used to manage blobs of chunks for logs.
type LogChunks struct {
	logger      *zap.SugaredLogger
	config      *entities.Config
	blobStorage entities.BlobStorage
}

// New returns a new instance of LogChunks.
func New(c *entities.Config, blobStorage entities.BlobStorage) *LogChunks {
	logger := c.RootLogger.Named("LogChunks").Sugar()

	return &LogChunks{
		logger:      logger,
		config:      c,
		blobStorage: blobStorage,
	}
}

const (
	fileNameExtension = "log"
	nameDateFormat    = "20060102_1504" // yyyymmdd_hhmm
)

func generateLogChunkName() string {
	// Get the current time
	now := time.Now()
	timestamp := now.Format(nameDateFormat)
	randomUUID := uuid.New()

	// Combine the timestamp and UUID to create the filename
	filename := fmt.Sprintf("%s_%s.%s", timestamp, randomUUID, fileNameExtension)

	return filename
}

// SaveChunk saves the chunk passed in and returns the filename.
func (ls *LogChunks) SaveChunk(logChunk []byte) (string, error) {
	fileName := generateLogChunkName()

	log := ls.logger.Named("SaveChunk").With("file_name", fileName)

	log.Debug("saving chunk")

	err := ls.blobStorage.UploadFile(fileName, logChunk)
	if err != nil {
		return "", fmt.Errorf("error saving chunk to blob storage: %w", err)
	}

	log.Info("file stored")

	return fileName, nil
}

// GetChunk loads the chunk for a given filename.
func (ls *LogChunks) GetChunk(chunkPath string) ([]byte, error) {
	log := ls.logger.Named("GetChunk").With("chunk_path", chunkPath)

	log.Debug("retrieving chunk")

	chunk, err := ls.blobStorage.DownloadFile(chunkPath)
	if err != nil {
		return nil, fmt.Errorf("error fetching chunk from blob storage: %w", err)
	}

	log.Info("file retrieved")

	return chunk, nil
}

// DeleteChunk deletes a chunk with a given path.
func (ls *LogChunks) DeleteChunk(chunkPath string) error {
	return ls.blobStorage.DeleteFile(chunkPath)
}
