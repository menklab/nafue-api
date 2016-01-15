package services

import (
	"sparticus/repositories"
	"sparticus/models/display"
	"sparticus/models/domain"
	"github.com/nu7hatch/gouuid"
	"sparticus/config"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/aws"
	"time"
)

type IFileService interface {
	AddFile(*display.FileDisplay) (error)
}

type FileService struct {
	fileRepository repositories.IFileRepository
}

func NewFileService(fileRepository repositories.IFileRepository) *FileService {
	return &FileService{fileRepository}
}

func (self *FileService) AddFile(fileDisplay *display.FileDisplay) (error) {

	// generate random uuid
	s3u, err := uuid.NewV4()
	if err != nil {
		return err
	}
	shortUrl, err := uuid.NewV4()
	if err != nil {
		return err
	}

	// create domain model from display
	 file := models.File{
		S3Path: config.S3Key + "/" + s3u.String(),
		ShortUrl: shortUrl.String(),
		TTL: fileDisplay.TTL,
	}

	// add file to db
	err = self.fileRepository.AddFile(&file)
	if err != nil {
		return err
	}
	fileDisplay.Id = file.Id


	// create put request on s3
	req, _ := GetS3Service().PutObjectRequest(&s3.PutObjectInput{
		Bucket: aws.String(config.S3Bucket),
		Key: aws.String(config.S3Key + "/" + s3u.String()),

	})
	url, err := req.Presign(15 * time.Minute)
	fileDisplay.UploadUrl = url

	return nil
}
