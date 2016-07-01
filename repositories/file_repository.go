package repositories

import (
	"log"
	"github.com/menkveldj/nafue-api/models"
	"time"
	"github.com/jmoiron/sqlx"
)

type IFileRepository interface {
	GetFile(*models.FileHeader) error
	AddFileHeader(*models.FileHeader) error
	DeleteFile(*models.FileHeader) error
}

type FileRepository struct {
	database *sqlx.DB
}

func NewFileRepository(d *sqlx.DB) *FileRepository {
	return &FileRepository{d}
}

func (self *FileRepository) GetFile(file *models.FileHeader) error {
	err := self.database.QueryRow(`
	SELECT id, _salt, hmac, ttl, created FROM files WHERE shortUrl = ?
	`, file.ShortUrl).Scan(&file.Id, &file.Salt, &file.Hmac, &file.TTL, &file.Created)
	if err != nil {
		log.Println("---ERROR---", err.Error())
		return err
	}

	return nil
}

func (self *FileRepository) AddFileHeader(file *models.FileHeader) error {
	now := time.Now()
	result, err := self.database.Exec(`
	INSERT INTO files
	(ttl, shortURL, created, _salt, hmac) VALUES (?,?,?,?,?,?,?)
	`, file.TTL, file.ShortUrl, now, file.Salt, file.Hmac)
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
