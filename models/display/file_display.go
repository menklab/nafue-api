package display

import "fmt"

type FileDisplay struct {
	TTL         int    `json:"ttl,omitempty"`
	ShortUrl    string `json:"shortUrl,omitempty"`
	UploadUrl   string `json:"uploadUrl,omitempty"`
	DownloadUrl string `json:"downloadUrl,omitempty"`
	IV          string `json:"iv" binding:"required"`
	Salt        string `json:"salt" binding:"required"`
	AData       string `json:"aData" binding:"required"`
}

func (self *FileDisplay) ToString() string {
	return fmt.Sprintf(
		"{Id: %v, UploadUrl: %v, TTL: %v, ShortURL: %v, IV: %v, Salt: %v, AData: %v}",
		self.UploadUrl,
		self.TTL,
		self.ShortUrl,
		self.IV,
		self.Salt,
		self.AData,
	)
}
