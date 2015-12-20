package services

import (
	"sparticus/repositories"
)

type IFileService interface {
	AddFile() (int, error)
}

type FileService struct {
	fileRepository repositories.IFileRepository
}

func NewFileService(fileRepository repositories.IFileRepository) *FileService {
	return &FileService{fileRepository}
}

func (self *FileService) AddFile() (int, error) {
	id, err := self.fileRepository.AddFile()
	if err != nil {
		return 0, err
	}

	return id, nil
}
