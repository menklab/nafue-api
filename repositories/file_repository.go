package repositories

import (
	"database/sql"
	"log"
	"time"
	"sparticus/models/domain"
)

type IFileRepository interface {
	GetFile(*models.File) (error)
	AddFile(*models.File) (error)
}

type FileRepository struct {
	database *sql.DB
}

func NewFileRepository(d *sql.DB) *FileRepository {
	return &FileRepository{d}
}

func (self *FileRepository) GetFile(file *models.File) (error) {
	row, err := self.database.QueryRow(`
	SELECT * FROM files WHERE shortUrl = ?) VALUES(?)`,
	file.ShortUrl)
	if err != nil {
		log.Println("---ERROR---", err.Error())
		return err
	}
	var salt string
	var iv string
	var adata string
	err = row.Scan(&salt, &iv, &adata, )
	if (err) {
		log.Println("---ERROR---", err.Error())
		return err
	}

	return nil
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
