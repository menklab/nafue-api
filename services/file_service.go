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
	GetFile(string) (*models.FileDisplay, error)
	AddFile(*models.FileHeader) (*models.FileDisplay, error)
}

type FileService struct {
	fileRepository           repositories.IFileRepository
	basicAnalyticsRepository repositories.IBasicAnalyticsRepository
}

func NewFileService(fileRepository repositories.IFileRepository, basicAnalyticsRepository repositories.IBasicAnalyticsRepository) *FileService {
	return &FileService{fileRepository, basicAnalyticsRepository}
}

func (self *FileService) GetFile(shortUrl string) (*models.FileDisplay, error) {

	// get file from db
	fileDisplay, err := self.fileRepository.GetFile(shortUrl)
	if err != nil {
		return nil, err
	}

	// now that we have file delete it from db
	self.fileRepository.DeleteFile(fileDisplay.FileHeader.Id)

	// verify that file isn't to old
	elapsed := int64(time.Now().Sub(fileDisplay.FileHeader.Created).Nanoseconds())
	if elapsed > fileDisplay.FileHeader.TTL {
		// to old delete file
		fmt.Println("file to old, delete from s3!")
		_, err := GetS3Service().DeleteObject(&s3.DeleteObjectInput{
			Bucket: aws.String(config.S3Bucket),
			//Key:    aws.String(config.S3Key + "/" + file.S3Path),
		})
		if err != nil {
			fmt.Println("Service Error: ", err.Error())
		}
		return nil, errors.New("File has expired")
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

	return fileDisplay, nil
}

func (self *FileService) AddFile(fileHeader *models.FileHeader) (*models.FileDisplay, error) {

	// get short url
	shortUrl, err := utility.GenerateRandomString(32)
	if err != nil {
		return nil, err
	}
	fileHeader.ShortUrl = shortUrl
	//fileHeader.TTL =  (1 * 60 * 60 * 24) // 24h in seconds
	fileHeader.TTL = int64(time.Minute) * 15

	// add file to db
	err = self.fileRepository.AddFileHeader(fileHeader)
	if err != nil {
		return nil, err
	}

	// calc num and chunk
	chunkSize := config.ChunkSize * 1024 * 1024 // convert to byte
	tChunks := int64(math.Ceil(float64(fileHeader.Size / chunkSize)))
	lastChunkSize := fileHeader.Size - (chunkSize * (tChunks-1))
	chunks := make([]models.FileChunk, tChunks)
	c := make(chan models.FileChunk)
	e := make(chan error)
	// spin off
	for i, chunk := range chunks {
		chunk.Order = i
		chunk.FileId = fileHeader.Id
		if i == (len(chunks) -1) {
			chunk.Size = lastChunkSize // last chunk is likely smaller
		} else {
			chunk.Size = chunkSize
		}
		go self.chunkIt(chunk, c, e)
	}
	// wait till all return
	for j := 0; j < int(tChunks); j++ {
		select {
		case err := <-e:
			return nil, err
		case mc := <-c:
			chunks[mc.Order] = mc
			break;
		}
	}

	fileDisplay := models.FileDisplay{
		FileHeader: models.FileHeader{
			ShortUrl: fileHeader.ShortUrl,
		},
		FileChunks: chunks,
	}


	self.basicAnalyticsRepository.IncrementFileCount()

	return &fileDisplay, nil
}

func (self *FileService) chunkIt(chunk models.FileChunk, c chan models.FileChunk, e chan error) {
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
	chunk.S3Path = ranName

	// save chunk to db
	err = self.fileRepository.AddFileChunk(&chunk)
	if err != nil {
		e <- err
		return
	}

	c <- chunk
	return
}
