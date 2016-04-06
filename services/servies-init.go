package services

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"nafue-api/config"
	"nafue-api/repositories"
)

var (
	fileService           IFileService
	s3Service             *s3.S3
	paymentService        IPaymentService
	basicAnalyticsService IBasicAnalyticsService
)

func Init() {
	fileService = NewFileService(repositories.GetFileRepository(), repositories.GetBasicAnalyticsRepository())
	s3Service = s3.New(session.New(&aws.Config{Region: aws.String(config.S3Location)}))
	paymentService = NewPaymentService()
	basicAnalyticsService = NewBasicAnalyticsService(repositories.GetBasicAnalyticsRepository())
}

// Public Getter
func GetFileService() IFileService {
	return fileService
}
func GetS3Service() *s3.S3 {
	return s3Service
}
func GetPaymentService() IPaymentService {
	return paymentService
}
func GetBasicAnalyticsService() IBasicAnalyticsService {
	return basicAnalyticsService
}
