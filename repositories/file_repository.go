package repositories

import (
	"database/sql"
	"log"
	"time"
	"sparticus/models/domain"
)

type IFileRepository interface {
	AddFile(*models.File) (error)
}

type FileRepository struct {
	database *sql.DB
}

func NewFileRepository(d *sql.DB) *FileRepository {
	return &FileRepository{d}
}

func (self *FileRepository) AddFile(file *models.File) (error) {
	now := time.Now()
	result, err := self.database.Exec(`
	INSERT INTO files
	(s3Path, ttl, shortURL, created, uploadUrl, iV, _salt, aData) VALUES (?,?,?,?,?,?,?,?)
	`, file.S3Path, file.TTL, file.ShortUrl, now, file.UploadUrl, file.IV, file.Salt, file.AData)
	if err != nil {
		log.Println("---ERROR---", err.Error())
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	file.Id = int(id)

	return nil
}
