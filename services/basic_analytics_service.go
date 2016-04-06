package services

import (
	"log"
	"nafue-api/models/display"
	"nafue-api/models/domain"
	"nafue-api/repositories"
)

type IBasicAnalyticsService interface {
	GetFileCount(*display.BasicAnalyticsDisplay) error
	IncrementFileCount(*display.BasicAnalyticsDisplay) error
}

type BasicAnalyticsService struct {
	basicAnalyticsRepository repositories.IBasicAnalyticsRepository
}

func NewBasicAnalyticsService(basicAnalyticsRepository repositories.IBasicAnalyticsRepository) *BasicAnalyticsService {
	return &BasicAnalyticsService{basicAnalyticsRepository}
}

func (self *BasicAnalyticsService) GetFileCount(basicAnalyticsDisplay *display.BasicAnalyticsDisplay) error {

	var basicAnalytics models.BasicAnalytics

	err := self.basicAnalyticsRepository.GetFileCount(&basicAnalytics)
	if err != nil {
		log.Println("error getting file count: ", err.Error())
		return err
	}
	basicAnalyticsDisplay.FileCount = basicAnalytics.FileCount

	return nil
}

func (self *BasicAnalyticsService) IncrementFileCount(basicAnalyticsDisplay *display.BasicAnalyticsDisplay) error {

	err := self.basicAnalyticsRepository.IncrementFileCount()
	if err != nil {
		log.Println("error incrementing file count: ", err.Error())
		return err
	}

	return nil
}
