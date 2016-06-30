package models

import "time"

type FileHeader struct {
	Id        int       `json:"id"`
	S3Path    string    `json:"s3Path,omitempty"`
	TTL       int       `json:"ttl,omitempty"`
	Created   time.Time `json:"created,omitempty"`
	ShortUrl  string    `json:"shortUrl,omitempty"`
	UploadUrl string    `json:"uploadUrl,omitempty"`
	DownloadUrl string `json:"downloadUrl,omitempty"`
	Salt      []byte    `json:"salt" binding:"required"`
	Hmac      []byte    `json:"hmac" binding:"required"`
}
