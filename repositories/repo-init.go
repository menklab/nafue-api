package repositories

import (
	"database/sql"
	"errors"
)

var (
	fileRepository           IFileRepository
	basicAnalyticsRepository IBasicAnalyticsRepository
)

func Init(database *sql.DB) {
	fileRepository = NewFileRepository(database)
	basicAnalyticsRepository = NewBasicAnalyticsRepository(database)
}

// Public Getters
func GetFileRepository() IFileRepository {
	return fileRepository
}

func GetBasicAnalyticsRepository() IBasicAnalyticsRepository {
	return basicAnalyticsRepository
}

//Helper Methods
func checkErrorType(err error, noRowsErrMessage, systemErrMessage string) error {
	var e error
	if err == sql.ErrNoRows {
		e = errors.New(noRowsErrMessage)
	} else {
		e = errors.New(systemErrMessage)
	}
	return e
}
