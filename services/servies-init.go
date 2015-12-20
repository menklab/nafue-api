package services

import (
	"sparticus/repositories"
)

var (
	fileService IFileService
)

func Init() {
	fileService = NewFileService(repositories.GetFileRepository())
}

// Public Getter
func GetFileService() IFileService {
	return fileService
}
