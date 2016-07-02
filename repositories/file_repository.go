package repositories

import (
	"log"
	"github.com/menkveldj/nafue-api/models"
	"github.com/jmoiron/sqlx"
	"time"
)

type IFileRepository interface {
	GetFile(string) (*models.FileDisplay, error)
	AddFileHeader(*models.FileHeader) error
	DeleteFile(int64) error
	AddFileChunk(fileChunk *models.FileChunk) error
}

type FileRepository struct {
	database *sqlx.DB
}

func NewFileRepository(d *sqlx.DB) *FileRepository {
	return &FileRepository{d}
}

func (self *FileRepository) GetFile(shortUrl string) (*models.FileDisplay, error) {

	// get file header
	var fileHeader models.FileHeader
	err := self.database.Get(&fileHeader, `SELECT * FROM files WHERE shortUrl = ?`,
		shortUrl)
	if err != nil {
		log.Println("DB ERROR", err.Error())
		return nil, err
	}

	// get file chunks
	var fileChunks []models.FileChunk
	err = self.database.Select(&fileChunks, `SELECT * from file_chunks WHERE fileId = ?`, fileHeader.Id)
	if err != nil {
		log.Println("DB ERROR", err.Error())
		return nil, err
	}

	fileDisplay := models.FileDisplay{
		FileHeader: fileHeader,
		FileChunks:fileChunks,
	}

	return &fileDisplay, nil
}

func (self *FileRepository) AddFileHeader(fileHeader *models.FileHeader) error {

	fileHeader.Created = time.Now().UTC()

	result, err := self.database.NamedExec(`
	INSERT INTO files
	(ttl, shortURL, _salt, hmac, created) VALUES (:ttl, :shortUrl, :_salt, :hmac, :created)
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


func (self *FileRepository) DeleteFile(fileId int64) error {
	_, err := self.database.Exec(`
	DELETE FROM files WHERE id = ?
	`, fileId)
	if err != nil {
		log.Println("DB ERROR", err.Error())
		return err
	}

	_, err = self.database.Exec(`
	DELETE FROM file_chunks WHERE fileId = ?
	`, fileId)
	if err != nil {
		log.Println("DB ERROR", err.Error())
		return err
	}

	return nil
}


func (self *FileRepository) AddFileChunk(fileChunk *models.FileChunk) error {

	fileChunk.Created = time.Now().UTC()

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