package services

import (
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/menkveldj/nafue-api/config"
	"github.com/menkveldj/nafue-api/models"
	"github.com/menkveldj/nafue-api/repositories"
	"time"
	"github.com/menkveldj/nafue-api/utility"
	"math"
)

type IFileService interface {
	GetFile(*models.FileHeader) error
	AddFile(*models.FileHeader) error
}

type FileService struct {
	fileRepository           repositories.IFileRepository
	basicAnalyticsRepository repositories.IBasicAnalyticsRepository
}

func NewFileService(fileRepository repositories.IFileRepository, basicAnalyticsRepository repositories.IBasicAnalyticsRepository) *FileService {
	return &FileService{fileRepository, basicAnalyticsRepository}
}

func (self *FileService) GetFile(fileDisplay *models.FileHeader) error {

	// make model from display
	file := models.FileHeader{
		ShortUrl: fileDisplay.ShortUrl,
	}

	// get file from db
	err := self.fileRepository.GetFile(&file)
	if err != nil {
		return err
	}

	// now that we have file delete it from db
	self.fileRepository.DeleteFile(&file)

	// verify that file isn't to old
	elapsed := int(time.Now().Sub(file.Created).Seconds())
	if elapsed > file.TTL {
		// to old delete file
		fmt.Println("file to old, delete from s3!")
		_, err := GetS3Service().DeleteObject(&s3.DeleteObjectInput{
			Bucket: aws.String(config.S3Bucket),
			//Key:    aws.String(config.S3Key + "/" + file.S3Path),
		})
		if err != nil {
			fmt.Println("---ERROR---", err.Error())
		}
		return errors.New("File has expired")
	}

	// create get request
	//req, _ := GetS3Service().GetObjectRequest(&s3.GetObjectInput{
	//	Bucket: aws.String(config.S3Bucket),
	//	Key:    aws.String(config.S3Key + "/" + file.S3Path),
	//})

	//url, err := req.Presign(15 * time.Minute)
	//if err != nil {
	//	log.Println("--ERROR---", err.Error())
	//	return err
	//}

	// add needed data to display
	//fileDisplay.DownloadUrl = url
	fileDisplay.Salt = file.Salt
	fileDisplay.Hmac = file.Hmac

	return nil
}

func (self *FileService) AddFile(fileHeader *models.FileHeader) error {

	shortUrl, err := utility.GenerateRandomString(32)
	if err != nil {
		return err
	}

	fileHeader.ShortUrl = shortUrl

	// add file to db
	//err = self.fileRepository.AddFileHeader(fileHeader)
	//if err != nil {
	//	return err
	//}

	chunkSize := config.ChunkSize * 1024 * 1024 // convert to mb

	// calc num and chunk
	tChunks := int64(math.Ceil(float64(fileHeader.Size / int(chunkSize))))
	chunks := make([]models.FileHeaderChunk, int(tChunks))
	c := make(chan models.FileHeaderChunk)
	e := make(chan error)
	// fan off
	for i, chunk := range chunks {
		chunk.Order = i
		go chunkIt(chunk, c, e)
	}
	// fan in
	for j := 0; j < int(tChunks); j++ {
		select {
		case err := <-e:
			fmt.Println("c error: ", err)
			return err
		case mc := <-c:
			fmt.Println("url: ", mc.UploadUrl)
			break;
		}
	}

	fmt.Println("chunks: ", chunks)
	// create domain model from display
	//file := models.FileHeader{
	//	S3Path:    s3u.String(),
	//ShortUrl:  shortUrl.String(),
	//TTL:       (1 * 60 * 60 * 24), // 24h in seconds
	//Salt:      fileDisplay.Salt,
	//Hmac:      fileDisplay.Hmac,
	//UploadUrl: url,
	//}

	// add upload url to display
	//fileDisplay.UploadUrl = file.UploadUrl
	//fileDisplay.ShortUrl = file.ShortUrls


	//self.basicAnalyticsRepository.IncrementFileCount()

	return nil
}

func chunkIt(chunk models.FileHeaderChunk, c chan models.FileHeaderChunk, e chan error) {
	// get random filename
	ranName, err := utility.GenerateRandomString(32)
	if err != nil {
		e <- err
		return
	}

	// create put request on s3
	req, _ := GetS3Service().PutObjectRequest(&s3.PutObjectInput{
		Bucket:      aws.String(config.S3Bucket),
		Key:         aws.String(config.S3Key + "/" + ranName),
		ContentType: aws.String("text/plain;charset=UTF-8"),
	})
	url, err := req.Presign(time.Duration(config.PresignLimit) * time.Hour)
	if err != nil {
		e <- err
		return
	}
	chunk.UploadUrl = url

	// save chunk to db

	c <- chunk
	return
}
