package models

import "time"

type File struct {
	Id        int       `json:"id"`
	S3Path    string    `json:"s3Path,omitempty"`
	TTL       int       `json:"ttl,omitempty"`
	Created   time.Time `json:"created,omitempty"`
	ShortUrl  string    `json:"shortUrl,omitempty"`
	UploadUrl string    `json:"uploadUrl,omitempty"`
	IV        string    `json:"iv" binding:"required"`
	Salt      string    `json:"salt" binding:"required"`
	AData     string    `json:"aData" binding:"required"`
}
