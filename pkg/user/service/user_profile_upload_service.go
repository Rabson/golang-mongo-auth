package service

import (
	"fmt"
	s3fileupload "golang-mongo-auth/pkg/libs/s3_file_upload"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// UploadToS3 uploads a file to S3 and returns the uploaded file URL
func uploadProfile(filePath string, userId string) (string, error) {
	// Mock implementation for S3 upload
	uploader := s3fileupload.GetS3Uploader()
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	input := &s3manager.UploadInput{
		Bucket: aws.String("your-bucket-name"),      // Replace with your bucket name
		Key:    aws.String(userId + "/" + filePath), // Construct the key using userId and filePath
		Body:   file,                                // Pass the opened file as the body
	}
	uploadedURL, err := uploader.Upload(input)
	if err != nil {
		return "", fmt.Errorf("failed to upload file to S3: %w", err)
	}
	return uploadedURL.Location, nil
}
