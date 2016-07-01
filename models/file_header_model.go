package models

import "time"

type FileHeader struct {
	Id       int       `json:"id" db:"id"`
	TTL      int       `json:"ttl,omitempty" db:"ttl"`
	Created  time.Time `json:"created,omitempty" db:"created"`
	ShortUrl string    `json:"shortUrl,omitempty" db:"shortUrl"`
	Salt     string    `json:"salt" binding:"required" db:"_salt"`
	Size	 int	   `json:"size,omitempty" binding:"required"`
	Hmac     string    `json:"hmac" binding:"required" db:"hmac"`
}

type FileHeaderChunk struct {
	Id          int `json:"id" db:"id"`
	S3Path      string    `json:"s3Path,omitempty" db:"s3Path"`
	Size      string    `json:"size,omitempty" db:"_size"`
	UploadUrl   string    `json:"uploadUrl,omitempty"`
	DownloadUrl string `json:"downloadUrl,omitempty"`
	Order	int `json:"order" db:"_order"`
}

