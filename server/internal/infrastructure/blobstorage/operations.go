package blobstorage

import (
	"bytes"
	"context"
	"fmt"

	"github.com/minio/minio-go/v7"
)

// DownloadFile downloads a file from blob storage.
func (bs BlobStorage) DownloadFile(path string) ([]byte, error) {
	ctx := context.TODO()
	log := bs.logger.Named("DownloadFile")

	obj, err := bs.storageClient.GetObject(ctx, bs.bucket, path, minio.GetObjectOptions{})
	if err != nil {
		return nil, fmt.Errorf("cannot get object from storage: %w", err)
	}

	objInfo, err := obj.Stat()
	if err != nil {
		return nil, fmt.Errorf("cannot get object info: %w", err)
	}

	log = log.With("size", objInfo.Size)
	log.Debug("fetched object from blob storage")

	buf := new(bytes.Buffer)

	_, err = buf.ReadFrom(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot read file: %w", err)
	}

	return buf.Bytes(), nil
}

// UploadFile uploads a file to blob storage.
func (bs BlobStorage) UploadFile(path string, contents []byte) error {
	ctx := context.TODO()
	log := bs.logger.Named("UploadFile")

	contentsReader := bytes.NewReader(contents)

	obj, err := bs.storageClient.PutObject(ctx, bs.bucket, path, contentsReader, int64(len(contents)), minio.PutObjectOptions{})
	if err != nil {
		return fmt.Errorf("error storing object: %w", err)
	}

	log = log.With("object_key", obj.Key)
	log.Info("stored object in blob storage")

	return nil
}
