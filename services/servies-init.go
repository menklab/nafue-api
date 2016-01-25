package services

import (
	"nafue/repositories"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"nafue/config"
)

var (
	fileService IFileService
	s3Service *s3.S3
)

func Init() {
	fileService = NewFileService(repositories.GetFileRepository())
	s3Service = s3.New(session.New(&aws.Config{Region: aws.String(config.S3Location)}))

}

// Public Getter
func GetFileService() IFileService {
	return fileService
}
func GetS3Service() *s3.S3 {
	return s3Service
}
