package display
import "fmt"

type FileDisplay struct {
	Id       int    `json:"id"`
	TTL      int    `json:"ttl"`
	ShortUrl string `json:"shortUrl"`
	UploadUrl string `json:"uploadUrl"`
}

func (self *FileDisplay) ToString() string {
	return fmt.Sprintf("{Id: %v, UploadUrl: %v, TTL: %v, ShortURL: %v", self.Id, self.UploadUrl, self.TTL, self.ShortUrl)
}