package display

import "fmt"

type FileHeaderDisplay struct {
	TTL         int    `json:"ttl,omitempty"`
	ShortUrl    string `json:"shortUrl,omitempty"`
	UploadUrl   string `json:"uploadUrl,omitempty"`
	DownloadUrl string `json:"downloadUrl,omitempty"`
	Salt        []byte `json:"salt" binding:"required"`
	FileSize    int64    `json:"fileSize" binding:"required"`
	Hmac        []byte    `json:"hmac" binding:"required"`
	MD5Checksum []byte    `json:"md5Checksum,omitempty" binding:"required"`
}

func (self *FileHeaderDisplay) ToString() string {
	return fmt.Sprintf(
		"{UploadUrl: %v, DownloadUrl: %v, TTL: %v, ShortURL: %v, Salt: %v, FileSize: %v}",
		self.UploadUrl,
		self.DownloadUrl,
		self.TTL,
		self.ShortUrl,
		self.Salt,
		self.FileSize,
	)
}
