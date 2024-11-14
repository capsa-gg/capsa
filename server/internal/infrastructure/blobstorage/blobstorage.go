package blobstorage

import (
	"context"
	"fmt"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"go.uber.org/zap"

	"github.com/capsa-gg/capsa/server/internal/entities"
)

// BlobStorage is the blob storage utility to save and retrieve files in an abstracted way.
type BlobStorage struct {
	storageClient *minio.Client
	bucket        string
	logger        *zap.SugaredLogger
}

// New creates a new BlobStorage instance, and makes sure the bucket from entities.Config exists.
func New(c *entities.Config) (entities.BlobStorage, error) {
	log := c.RootLogger.Named("BlobStorage").Sugar()
	ctx := context.TODO()

	minioClient, err := minio.New(c.BlobStorageEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(c.BlobStorageKey, c.BlobStorageSecret, ""),
		Secure: true,
	})

	if err != nil {
		return BlobStorage{}, fmt.Errorf("error creating blob storage instance: %s", err)
	}

	log.Debugf("storage endpoint: %s", minioClient.EndpointURL())

	if err = ensureBucketExists(ctx, log, c, minioClient); err != nil {
		return BlobStorage{}, fmt.Errorf("bucket existence ensurance failed: %w", err)
	}

	blobStorage := BlobStorage{
		storageClient: minioClient,
		logger:        log,
		bucket:        c.BlobStorageBucket,
	}

	return blobStorage, nil
}

// ensureBucketExists checks if a bucket exists and creates it if it doesn't.
func ensureBucketExists(ctx context.Context, logger *zap.SugaredLogger, c *entities.Config, minioClient *minio.Client) error {
	bucketName := c.BlobStorageBucket
	log := logger.Named("ensureBucketExists").With("bucket_name", bucketName)

	exists, err := minioClient.BucketExists(ctx, bucketName)
	if err != nil {
		return fmt.Errorf("cannot check if bucket exists: %w", err)
	}

	if exists {
		log.Info("bucket exists")

		return nil
	}

	log.Warn("bucket does not exist, attempting to create")

	options := minio.MakeBucketOptions{
		Region:        c.BlobStorageRegion,
		ObjectLocking: false,
	}

	err = minioClient.MakeBucket(ctx, bucketName, options)

	if err != nil {
		return fmt.Errorf("error creating bucket: %w", err)
	}

	return nil
}
