package repositories

import (
	"log"
	"github.com/menkveldj/nafue-api/models"
	"github.com/jmoiron/sqlx"
)

type IFileRepository interface {
	GetFile(*models.FileHeader) error
	AddFileHeader(*models.FileHeader) error
	DeleteFile(*models.FileHeader) error
	AddFileChunk(fileChunk *models.FileChunk) error
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
		log.Println("DB ERROR", err.Error())
		return err
	}

	return nil
}

func (self *FileRepository) AddFileHeader(fileHeader *models.FileHeader) error {

	result, err := self.database.NamedExec(`
	INSERT INTO files
	(ttl, shortURL, _salt, hmac) VALUES (:ttl, :shortUrl, :_salt, :hmac)
	`, fileHeader)
	if err != nil {
		log.Println("DB ERROR", err.Error())
		return err
	}

	// add id to file
	id, _ := result.LastInsertId()
	fileHeader.Id = id

	return nil
}


func (self *FileRepository) DeleteFile(file *models.FileHeader) error {
	_, err := self.database.Exec(`
	DELETE FROM files WHERE id = ?
	`, file.Id)
	if err != nil {
		log.Println("DB ERROR", err.Error())
		return err
	}

	return nil
}


func (self *FileRepository) AddFileChunk(fileChunk *models.FileChunk) error {

	result, err := self.database.NamedExec(`
	INSERT INTO file_chunks
	(fileId, s3Path, _size, _order) VALUES (:fileId, :s3Path, :_size, :_order)
	`, fileChunk)
	if err != nil {
		log.Println("DB ERROR", err.Error())
		return err
	}

	// add id to file
	id, _ := result.LastInsertId()
	fileChunk.Id = id

	return nil
}