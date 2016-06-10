package repositories

import (
	"database/sql"
	"log"
	"github.com/menkveldj/nafue-api/models/domain"
)

type IBasicAnalyticsRepository interface {
	IncrementFileCount() error
	GetFileCount(*models.BasicAnalytics) error
}

type BasicAnalyticsRepository struct {
	database    *sql.DB
	fileCountId int
}

func NewBasicAnalyticsRepository(d *sql.DB) *BasicAnalyticsRepository {
	// verify file count exists and cache value for later
	var fileCountId int
	err := d.QueryRow(`SELECT id FROM basic_analytics WHERE field_name='uploadedFile'`).Scan(&fileCountId)
	// if error create row
	if err != nil {
		log.Println("---ERROR---", err.Error())
		// row doesn't exist so create row
		result, err := d.Exec(`INSERT INTO basic_analytics (field_name, int_value) VALUES('uploadedFile', '0')`)
		if err != nil {
			log.Println("---ERROR---", err.Error())
			return nil
		}
		id, err := result.LastInsertId()
		fileCountId = int(id)
	}
	return &BasicAnalyticsRepository{d, fileCountId}
}

func (self *BasicAnalyticsRepository) IncrementFileCount() error {
	_, err := self.database.Exec(`UPDATE basic_analytics SET int_value = int_value + 1 WHERE id = ?`, self.fileCountId)
	if err != nil {
		log.Println("---ERROR---", err.Error())
		return err
	}
	return nil
}

func (self *BasicAnalyticsRepository) GetFileCount(basicAnalyticsModel *models.BasicAnalytics) error {
	err := self.database.QueryRow(`
	SELECT int_value FROM basic_analytics WHERE id = ?
	`, self.fileCountId).Scan(&basicAnalyticsModel.FileCount)
	if err != nil {
		log.Println("---ERROR---", err.Error())
		return err
	}

	return nil
}
