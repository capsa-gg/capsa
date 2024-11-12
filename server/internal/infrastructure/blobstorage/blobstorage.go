package blobstorage

import (
	"context"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"go.uber.org/zap"

	"github.com/capsa-gg/capsa/server/internal/entities"
)

// BlobStorage is the blob storage utility to save and retrieve files in an abstracted way.
type BlobStorage struct {
	s3Client *s3.Client
	bucket   string
	logger   *zap.SugaredLogger
}

// New creates a new BlobStorage instance, and makes sure the bucket from entities.Config exists.
func New(c *entities.Config) (entities.BlobStorage, error) {
	log := c.RootLogger.Named("BlobStorage").Sugar()
	ctx := context.TODO()

	creds := credentials.NewStaticCredentialsProvider(c.BlobStorageKey, c.BlobStorageSecret, "")

	customResolver := aws.EndpointResolverWithOptionsFunc(func(_, _ string, _ ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL: c.BlobStorageEndpoint,
		}, nil
	})

	awsConfig, err := config.LoadDefaultConfig(ctx,
		config.WithRegion(c.BlobStorageRegion),
		config.WithCredentialsProvider(creds),
		config.WithEndpointResolverWithOptions(customResolver),
	)

	if err != nil {
		return BlobStorage{}, fmt.Errorf("error creating aws config: %s", err)
	}

	s3Client := s3.NewFromConfig(awsConfig)

	if err = ensureBucketExists(ctx, log, c, s3Client); err != nil {
		return BlobStorage{}, fmt.Errorf("bucket existence ensurance failed: %w", err)
	}

	blobStorage := BlobStorage{
		s3Client: s3Client,
		logger:   log,
		bucket:   c.BlobStorageBucket,
	}

	return blobStorage, nil
}

// ensureBucketExists checks if a bucket exists and creates it if it doesn't.
func ensureBucketExists(ctx context.Context, logger *zap.SugaredLogger, c *entities.Config, s3Client *s3.Client) error {
	log := logger.Named("ensureBucketExists")
	bucketName := c.BlobStorageBucket

	// Check if the bucket exists
	_, err := s3Client.HeadBucket(ctx, &s3.HeadBucketInput{
		Bucket: aws.String(bucketName),
	})

	if err != nil {
		// If the error is because the bucket doesn't exist, create it
		if isNotFoundError(err) {
			_, err = s3Client.CreateBucket(ctx, &s3.CreateBucketInput{
				Bucket: aws.String(bucketName),
				CreateBucketConfiguration: &types.CreateBucketConfiguration{
					LocationConstraint: types.BucketLocationConstraint(c.BlobStorageRegion),
				},
			})

			if err != nil {
				return fmt.Errorf("error creating bucket: %s", err)
			}

			log.Infof("bucket %s created successfully", bucketName)
		} else {
			return fmt.Errorf("error checking bucket existence: %s", err)
		}
	} else {
		log.Infof("bucket %s already exists", bucketName)
	}

	return nil
}

// isNotFoundError checks if the error is a bucket not found error.
func isNotFoundError(err error) bool {
	var notFound *types.NotFound
	return errors.As(err, &notFound)
}
