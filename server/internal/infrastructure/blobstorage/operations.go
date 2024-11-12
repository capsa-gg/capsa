package blobstorage

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/capsa-gg/capsa/server/internal/entities"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

// DownloadFile downloads a file from blob storage.
func (bs BlobStorage) DownloadFile(path string) ([]byte, error) {
	ctx := context.TODO()
	log := bs.logger.Named("DownloadFile")

	result, err := bs.s3Client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(bs.bucket),
		Key:    aws.String(path),
	})

	if err != nil {
		var noKey *types.NoSuchKey

		if errors.As(err, &noKey) {
			log.Warnf("can't get object %s from bucket %s, no such key exists.", path, bs.bucket)

			return nil, entities.NewDomainError(entities.DomainErrorNotFound, fmt.Sprintf("file %s not found", path), err)
		}

		log.Errorf("can't get object %s from bucket %s: %s", path, bs.bucket, err)

		return nil, fmt.Errorf("error fetching file %s: %w", path, err)
	}

	defer func(Body io.ReadCloser) {
		errClose := Body.Close()
		if errClose != nil {
			log.Errorf("error closing body: %s", err)
		}
	}(result.Body)

	body, err := io.ReadAll(result.Body)

	if err != nil {
		log.Errorf("error reading body for file %s: %s", path, err)

		return nil, fmt.Errorf("error reading file %s body: %w", path, err)
	}

	return body, nil
}

// UploadFile uploads a file to blob storage.
func (bs BlobStorage) UploadFile(path string, contents []byte) error {
	ctx := context.TODO()
	log := bs.logger.Named("UploadFile")

	contentsReader := bytes.NewReader(contents)

	_, err := bs.s3Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bs.bucket),
		Key:    aws.String(path),
		Body:   contentsReader,
	})

	if err != nil {
		log.Errorf("error storing file %s: %s", path, err)

		return fmt.Errorf("error storing file %s: %w", path, err)
	}

	return nil
}
