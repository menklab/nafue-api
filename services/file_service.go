package services

import (
	"sparticus/repositories"
	"sparticus/models/display"
	"sparticus/models/domain"
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

	// create domain model from display
	 file := models.File{
		S3Path: fileDisplay.S3Path,
		ShortUrl: fileDisplay.ShortUrl,
		TTL: fileDisplay.TTL,
	}

	err := self.fileRepository.AddFile(&file)
	if err != nil {
		return err
	}

	fileDisplay.Id = file.Id

	return nil
}
