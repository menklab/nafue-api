package repositories

import (
	"database/sql"
	"log"
	"time"
)

type IFileRepository interface {
	AddFile() (int, error)
}

type FileRepository struct {
	database *sql.DB
}

func NewFileRepository(d *sql.DB) *FileRepository {
	return &FileRepository{d}
}

func (self *FileRepository) AddFile() (int, error) {
	result, err := self.database.Exec(`
	INSERT INTO files
	(s3Path, ttl, shortURL, created) VALUES (?,?,?,?)
	`, "http://www.google.com", 2, "abcd", time.Now())
	if err != nil {
		log.Println("---ERROR---", err.Error())
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}
