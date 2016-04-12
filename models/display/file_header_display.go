package display

import "fmt"

type FileHeaderDisplay struct {
	TTL         int    `json:"ttl,omitempty"`
	ShortUrl    string `json:"shortUrl,omitempty"`
	UploadUrl   string `json:"uploadUrl,omitempty"`
	DownloadUrl string `json:"downloadUrl,omitempty"`
	IV          []byte `json:"iv" binding:"required"`
	Salt        []byte `json:"salt" binding:"required"`
	AData       []byte `json:"aData" binding:"required"`
}

func (self *FileHeaderDisplay) ToString() string {
	return fmt.Sprintf(
		"{UploadUrl: %v, DownloadUrl: %v, TTL: %v, ShortURL: %v, IV: %v, Salt: %v, AData: %v}",
		self.UploadUrl,
		self.DownloadUrl,
		self.TTL,
		self.ShortUrl,
		self.IV,
		self.Salt,
		self.AData,
	)
}
