package models

import "time"

type FileHeader struct {
	Id        int       `json:"id"`
	S3Path    string    `json:"s3Path,omitempty"`
	TTL       int       `json:"ttl,omitempty"`
	Created   time.Time `json:"created,omitempty"`
	ShortUrl  string    `json:"shortUrl,omitempty"`
	UploadUrl string    `json:"uploadUrl,omitempty"`
	IV        []byte    `json:"iv" binding:"required"`
	Salt      []byte    `json:"salt" binding:"required"`
	AData     []byte    `json:"aData" binding:"required"`
}
