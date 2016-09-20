package services

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/menkveldj/nafue-api/config"
	"github.com/menkveldj/nafue-api/models"
	"github.com/menkveldj/nafue-api/repositories"
	"time"
	"github.com/menkveldj/nafue-api/utility"
	"log"
	"github.com/menkveldj/nafue-api/utility/errors"
	"fmt"
)

type IFileService interface {
	GetFileAsync(string) (*models.FileDisplay, error)
	AddFileAsync(*models.FileHeader) (*models.FileHeader, error)
}

type FileService struct {
	fileRepository           repositories.IFileRepository
	basicAnalyticsRepository repositories.IBasicAnalyticsRepository
}

func NewFileService(fileRepository repositories.IFileRepository, basicAnalyticsRepository repositories.IBasicAnalyticsRepository) *FileService {
	return &FileService{fileRepository, basicAnalyticsRepository}
}

func (self *FileService) GetFileAsync(shortUrl string) (*models.FileDisplay, error) {

	// get file from db
	fileDisplay, err := self.fileRepository.GetFile(shortUrl)
	if err != nil {
		return nil, err
	}

	// now that we have file delete it from db
	go self.fileRepository.DeleteFile(fileDisplay.FileHeader.Id)

	// verify that file isn't to old
	elapsed := time.Since(fileDisplay.FileHeader.Created)
	if int64(elapsed) > fileDisplay.FileHeader.TTL {
		for _, chunk := range fileDisplay.FileChunks {
			go self.deleteChunks(chunk.S3Path)
		}
		return nil, errors.New("File expired!")
	}

	// get download urls for chunk
	c := make(chan models.FileChunk)
	e := make(chan error)
	// spin off
	for _, chunk := range fileDisplay.FileChunks {
		go self.chunkDownloadLink(chunk, c, e)
	}
	// spin in
	for i := 0; i < len(fileDisplay.FileChunks); i++ {
		select {
		case err := <-e:
			return nil, err
			break
		case mc := <-c:
			fileDisplay.FileChunks[fmt.Sprint(mc.Order)] = mc
			break
		}
	}

	// don't need tty so remove it
	fileDisplay.FileHeader.TTL = 0

	return fileDisplay, nil
}

func (self *FileService) AddFileAsync(fileHeader *models.FileHeader) (*models.FileHeader, error) {

	// generate sync key
	asyncKey, err := utility.GenerateRandomString(32)
	if err != nil {
		return nil, err
	}

	fileHeader.AsyncKey = asyncKey
	fileHeader.TTL = int64(time.Minute) * 15

	// add file to db
	err = self.fileRepository.AddFileHeaderAsync(fileHeader)
	if err != nil {
		return nil, err
	}

	tChunks := fileHeader.ChunkCount
	chunks := make(map[string]models.FileChunk)
	c := make(chan models.FileChunk)
	e := make(chan error)
	// spin off
	for i := 0; i < int(tChunks); i++ {

		// create chunk
		chunk := models.FileChunk{
			FileId: fileHeader.Id,
			Order: i,

		}
		// set chunk and save it
		chunks[fmt.Sprint(i)] = chunk
		go self.chunkIt(chunk, c, e)
	}
	// wait till all return
	for j := 0; j < int(tChunks); j++ {
		select {
		case err := <-e:
			return nil, err
		case mc := <-c:
			chunks[fmt.Sprint(mc.Order)] = mc
			break;
		}
	}

	// return async key
	fileHeaderDisplay := models.FileHeader{
		AsyncKey: asyncKey,
	}

	self.basicAnalyticsRepository.IncrementFileCount()

	return &fileHeaderDisplay, nil
}

func (self *FileService) chunkDownloadLink(chunk models.FileChunk, c chan models.FileChunk, e chan error) {

	// create get request
	req, _ := GetS3Service().GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(config.S3Bucket),
		Key:    aws.String(config.S3Key + "/" + chunk.S3Path),
	})

	url, err := req.Presign(time.Duration(config.PresignLimit) * time.Hour)
	if err != nil {
		log.Println("--ERROR---", err.Error())
		e <- err
		return
	}
	chunk.DownloadUrl = url
	c <- chunk
	return
}

func (self *FileService) chunkIt(chunk models.FileChunk, c chan models.FileChunk, e chan error) {
	// get random filename
	ranName, err := utility.GenerateRandomString(32)
	if err != nil {
		e <- err
		return
	}
	chunk.S3Path = ranName

	// ** DO THIS IF SYNC
	// create put request on s3
	//req, _ := GetS3Service().PutObjectRequest(&s3.PutObjectInput{
	//	Bucket:      aws.String(config.S3Bucket),
	//	Key:         aws.String(config.S3Key + "/" + ranName),
	//	ContentType: aws.String("text/plain;charset=UTF-8"),
	//})
	//url, err := req.Presign(time.Duration(config.PresignLimit) * time.Hour)
	//if err != nil {
	//	e <- err
	//	return
	//}
	//chunk.UploadUrl = url


	// save chunk to db
	err = self.fileRepository.AddFileChunk(&chunk)
	if err != nil {
		e <- err
		return
	}

	c <- chunk
	return
}

func (self *FileService) deleteChunks(s3Key string) {
	_, err := GetS3Service().DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(config.S3Bucket),
		Key:    aws.String(config.S3Key + "/" + s3Key),
	})
	if err != nil {
		log.Println("Error deleting chunk on s3: " + err.Error())
	}
}