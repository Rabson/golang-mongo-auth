package fileManager

import (
	"fmt"
	"golang-mongo-auth/pkg/config"
	s3fileupload "golang-mongo-auth/pkg/libs/s3_file_upload"
	"mime/multipart"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func UploadFile(fileType FileType, file *multipart.FileHeader, fileName string) (string, error) {

	if file == nil {
		return "", fmt.Errorf("file is nil")
	}

	if fileName == "" {
		return "", fmt.Errorf("fileName is empty")
	}

	if s3FilePath[fileType] == "" {
		return "", fmt.Errorf("Invalid fileType")
	}
	region := config.GetS3Region()

	uploader := s3fileupload.GetS3Uploader(region, config.GetS3AccessKeyId(), config.GetS3SecretAccessKey(), config.GetS3Endpoint())
	fileData, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}
	defer fileData.Close()

	filePath := s3FilePath[fileType]

	input := &s3manager.UploadInput{
		Bucket:      S3BucketName(),                        // Replace with your bucket name
		Key:         aws.String(filePath + "/" + fileName), // Construct the key using userId and filePath
		Body:        fileData,
		ACL:         aws.String("public-read"), // Set the ACL to public-read
		ContentType: aws.String(file.Header.Get("Content-Type")),
		Metadata: map[string]*string{
			"Content-Type": aws.String(file.Header.Get("Content-Type")),
			"FileName":     aws.String(file.Filename),
		},
		StorageClass: aws.String("STANDARD"),
		// Add any other metadata you want to include
		// e.g., "x-amz-meta-my-key": aws.String("my-value"),
		// You can also set other options like ServerSideEncryption, etc.
		// depending on your requirements.
		// ServerSideEncryption: aws.String("AES256"), // Example for server-side encryption
		// ContentDisposition:   aws.String("inline"),  // Example for content disposition
		// ContentEncoding:     aws.String("gzip"),    // Example for content encoding
		// ContentLanguage:     aws.String("en-US"),  // Example for content language
		// CacheControl:        aws.String("max-age=3600"), // Example for cache control
		// Expires:             aws.Time(time.Now().Add(24 * time.Hour)), // Example for expiration
		// Tagging:             aws.String("key1=value1&key2=value2"), // Example for tagging
		// WebsiteRedirectLocation: aws.String("/redirect"), // Example for website redirect location
		// GrantRead:           aws.String("uri=http://acs.amazonaws.com/groups/global/AllUsers"), // Example for granting read access
		// GrantWrite:          aws.String("uri=http://acs.amazonaws.com/groups/global/AllUsers"), // Example for granting write access
		// GrantFullControl:    aws.String("uri=http://acs.amazonaws.com/groups/global/AllUsers"), // Example for granting full control
		// GrantReadACP:        aws.String("uri=http://acs.amazonaws.com/groups/global/AllUsers"), // Example for granting read ACP access
		// GrantWriteACP:       aws.String("uri=http://acs.amazonaws.com/groups/global/AllUsers"), // Example for granting write ACP access
		// GrantReadAcp:        aws.String("uri=http://acs.amazonaws.com/groups/global/AllUsers"), // Example for granting read ACP access
		// GrantWriteAcp:       aws.String("uri=http://acs.amazonaws.com/groups/global/AllUsers"), // Example for granting write ACP access
		// GrantFullControl:    aws.String("uri=http://acs.amazonaws.com/groups/global/AllUsers"), // Example for granting full control
		// GrantRead:           aws.String("uri=http://acs.amazonaws.com/groups/global/AllUsers"), // Example for granting read access
		// GrantWrite:          aws.String("uri=http://acs.amazonaws.com/groups/global/AllUsers"), // Example for granting write access
		// GrantReadACP:        aws.String("uri=http://acs.amazonaws.com/groups/global/AllUsers"), // Example for granting read ACP access
	}
	uploadedURL, err := uploader.Upload(input)
	if err != nil {
		return "", fmt.Errorf("failed to upload file to S3: %w", err)
	}
	return uploadedURL.Location, nil
}
