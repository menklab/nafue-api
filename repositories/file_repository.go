package repositories

import (
	"database/sql"
	"log"
	"nafue-api/models/domain"
	"time"
)

type IFileRepository interface {
	GetFile(*models.FileHeader) error
	AddFile(*models.FileHeader) error
	DeleteFile(*models.FileHeader) error
}

type FileRepository struct {
	database *sql.DB
}

func NewFileRepository(d *sql.DB) *FileRepository {
	return &FileRepository{d}
}

func (self *FileRepository) GetFile(file *models.FileHeader) error {
	err := self.database.QueryRow(`
	SELECT id, _salt, iv, aData, s3Path, ttl, created FROM files WHERE shortUrl = ?
	`, file.ShortUrl).Scan(&file.Id, &file.Salt, &file.IV, &file.AData, &file.S3Path, &file.TTL, &file.Created)
	if err != nil {
		log.Println("---ERROR---", err.Error())
		return err
	}

	return nil
}

func (self *FileRepository) AddFile(file *models.FileHeader) error {
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

func (self *FileRepository) DeleteFile(file *models.FileHeader) error {
	_, err := self.database.Exec(`
	DELETE FROM files WHERE id = ?
	`, file.Id)
	if err != nil {
		log.Println("---ERROR---", err.Error())
		return err
	}

	return nil
}
