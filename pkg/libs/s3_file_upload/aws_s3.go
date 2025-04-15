package s3fileupload

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var s3Uploader *s3manager.Uploader

/*
- InitS3Uploader initializes the S3 uploader with the provided AWS session.

- It creates a new session with the specified region and initializes the S3 uploader.

- If the session creation fails, it logs the error and exits the program.

- This function should be called once at the start of the application to set up the S3 uploader.

	uploader := utils.GetS3Uploader() // Assume you have a function to get an S3 uploader

	_, err = uploader.Upload(&s3manager.UploadInput{
	Bucket: aws.String("your-s3-bucket-name"),
	Key:    aws.String("uploads/" + file.Filename),
	Body:   fileContent,
	})

	if err != nil {
	c.JSON(500, gin.H{"error": "Failed to save image"})
	return
	}
*/
func InitS3Uploader(region string, accessKeyID string, secretAccessKey string, endpoint string) {
	sess, err := session.NewSession(&aws.Config{
		Region:   aws.String(region),
		Endpoint: aws.String(endpoint),
		Credentials: credentials.NewStaticCredentials(
			accessKeyID,
			secretAccessKey,
			"",
		),
	})
	if err != nil {
		log.Fatalf("Failed to create AWS session: %v", err)
	}

	s3Uploader = s3manager.NewUploader(sess)
}

func GetS3Uploader(region string, accessKeyID string, secretAccessKey string, endpoint string) *s3manager.Uploader {
	if s3Uploader == nil {
		InitS3Uploader(region, accessKeyID, secretAccessKey, endpoint)
	}
	return s3Uploader
}
