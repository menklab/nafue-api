package display
import "fmt"

type FileDisplay struct {
	Id       int    `json:"id"`
	S3Path   string `json:"s3Path"`
	TTL      int    `json:"ttl"`
	ShortUrl string `json:"shortUrl"`
}

func (self *FileDisplay) ToString() string {
	return fmt.Sprintf("{Id: %v, S3Path: %v, TTL: %v, ShortURL: %v", self.Id, self.S3Path, self.TTL, self.ShortUrl)
}