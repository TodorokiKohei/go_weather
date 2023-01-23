package main

import (
	"context"
	"io"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type ObjectClinet struct {
	Client *minio.Client
}

func (obj *ObjectClinet) Connect(endpoint string, accessKeyId string, secretAccessKey string) (err error) {
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyId, secretAccessKey, ""),
		Secure: false,
	})
	if err != nil {
		return err
	}
	obj.Client = client
	return nil
}

func (obj *ObjectClinet) Post(ctx context.Context, bucketName string, objectName string, reader io.Reader, size int64) (info minio.UploadInfo, err error) {
	uploadInfo, err := obj.Client.PutObject(
		ctx,
		bucketName,
		objectName,
		reader,
		size,
		minio.PutObjectOptions{ContentType: "application/json"},
	)
	return uploadInfo, err
}
