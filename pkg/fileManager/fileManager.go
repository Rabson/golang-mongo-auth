package fileManager

import (
	"golang-mongo-auth/pkg/config"

	"github.com/aws/aws-sdk-go/aws"
)

type FileType string

func S3BucketName() *string {
	return aws.String(config.GetS3RBucket())
}
func S3Region() *string {
	return aws.String(config.GetS3Region())
}

const subPath = "fileManager"

const (
	FILE_TYPE_PROFILE       FileType = "profile"
	FILE_TYPE_DOCUMENT      FileType = "document"
	FILE_TYPE_IMAGE         FileType = "image"
	FILE_TYPE_VIDEO         FileType = "video"
	FILE_TYPE_AUDIO         FileType = "audio"
	FILE_TYPE_OTHER         FileType = "other"
	FILE_TYPE_DOCUMENT_PDF  FileType = "pdf"
	FILE_TYPE_DOCUMENT_DOC  FileType = "doc"
	FILE_TYPE_DOCUMENT_DOCX FileType = "docx"
	FILE_TYPE_DOCUMENT_XLS  FileType = "xls"
	FILE_TYPE_DOCUMENT_XLSX FileType = "xlsx"
	FILE_TYPE_DOCUMENT_PPT  FileType = "ppt"
	FILE_TYPE_DOCUMENT_PPTX FileType = "pptx"
	FILE_TYPE_DOCUMENT_TXT  FileType = "txt"
	FILE_TYPE_DOCUMENT_CSV  FileType = "csv"
	FILE_TYPE_DOCUMENT_ZIP  FileType = "zip"
	FILE_TYPE_DOCUMENT_RAR  FileType = "rar"
	FILE_TYPE_DOCUMENT_HTML FileType = "html"
	FILE_TYPE_DOCUMENT_XML  FileType = "xml"
	FILE_TYPE_DOCUMENT_JSON FileType = "json"
	FILE_TYPE_DOCUMENT_MD   FileType = "md"
	FILE_TYPE_DOCUMENT_YAML FileType = "yaml"
)

var s3FilePath = map[FileType]string{
	FILE_TYPE_PROFILE:      subPath + "/profile",
	FILE_TYPE_DOCUMENT:     subPath + "/document",
	FILE_TYPE_IMAGE:        subPath + "/upload",
	FILE_TYPE_VIDEO:        subPath + "/video",
	FILE_TYPE_AUDIO:        subPath + "/audio",
	FILE_TYPE_OTHER:        subPath + "/other",
	FILE_TYPE_DOCUMENT_PDF: subPath + "/docs",
	FILE_TYPE_DOCUMENT_DOC: subPath + "/docs",
}
