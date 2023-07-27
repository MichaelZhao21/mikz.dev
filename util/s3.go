package util

import (
	"bytes"
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// UploadToS3 will upload a given buffer to
// the S3 bucket specified by path
func UploadToS3(buffer []byte, path string) error {
	// Create config
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return err
	}

	// Create S3 client
	client := s3.NewFromConfig(cfg)
	bucket := GetEnv("AWS_S3_BUCKET")

	// Input body
	input := &s3.PutObjectInput{
		Body:   bytes.NewReader(buffer),
		Bucket: aws.String(bucket),
		Key:    aws.String(path),
	}

	// Make call
	_, err = client.PutObject(context.TODO(), input)
	if err != nil {
		return err
	}

	return nil
}
