package s3fileupload

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var s3Uploader *s3manager.Uploader

// uploader := utils.GetS3Uploader() // Assume you have a function to get an S3 uploader
//
//	_, err = uploader.Upload(&s3manager.UploadInput{
//		Bucket: aws.String("your-s3-bucket-name"),
//		Key:    aws.String("uploads/" + file.Filename),
//		Body:   fileContent,
//	})
//
//	if err != nil {
//		c.JSON(500, gin.H{"error": "Failed to save image"})
//		return
//	}
func InitS3Uploader() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("your-region"), // Replace with your AWS region
	})
	if err != nil {
		log.Fatalf("Failed to create AWS session: %v", err)
	}

	s3Uploader = s3manager.NewUploader(sess)
}

func GetS3Uploader() *s3manager.Uploader {
	if s3Uploader == nil {
		InitS3Uploader()
	}
	return s3Uploader
}
