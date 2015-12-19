package repositories
import (
	"database/sql"
	"sparticus/domain/models"
	"time"
	"log"
)

type IFileRepository interface {
	File()
}

type FileRepository struct {
	database *sql.DB
}

func NewItemRepository(d *sql.DB) *FileRepository {
	return &FileRepository{d}
}

func (r *FileRepository) AddFile() ( int, error) {
	result, err := r.database.Exec(`
	INSERT INTO files
	(S3Path, TTL, ShortURL, Created) VALUES (?,?,?,?),
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