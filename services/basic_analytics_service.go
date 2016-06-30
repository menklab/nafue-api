package services

import (
	"log"
	"github.com/menkveldj/nafue-api/models"
	"github.com/menkveldj/nafue-api/repositories"
)

type IBasicAnalyticsService interface {
	GetFileCount(*models.BasicAnalytics) error
	IncrementFileCount(*models.BasicAnalytics) error
}

type BasicAnalyticsService struct {
	basicAnalyticsRepository repositories.IBasicAnalyticsRepository
}

func NewBasicAnalyticsService(basicAnalyticsRepository repositories.IBasicAnalyticsRepository) *BasicAnalyticsService {
	return &BasicAnalyticsService{basicAnalyticsRepository}
}

func (self *BasicAnalyticsService) GetFileCount(basicAnalyticsDisplay *models.BasicAnalytics) error {

	var basicAnalytics models.BasicAnalytics

	err := self.basicAnalyticsRepository.GetFileCount(&basicAnalytics)
	if err != nil {
		log.Println("error getting file count: ", err.Error())
		return err
	}
	basicAnalyticsDisplay.FileCount = basicAnalytics.FileCount

	return nil
}

func (self *BasicAnalyticsService) IncrementFileCount(basicAnalyticsDisplay *models.BasicAnalytics) error {

	err := self.basicAnalyticsRepository.IncrementFileCount()
	if err != nil {
		log.Println("error incrementing file count: ", err.Error())
		return err
	}

	return nil
}
