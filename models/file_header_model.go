package models

import "time"

type FileHeader struct {
	Id       int64       `json:"-" db:"id"`
	TTL      int64       `json:"ttl,omitempty" db:"ttl"`
	Created  time.Time `json:"-" db:"created"`
	ShortUrl string    `json:"shortUrl,omitempty" db:"shortUrl"`
	Salt     string    `json:"salt,omitempty" binding:"required" db:"_salt"`
	Size     int64           `json:"size,omitempty" binding:"required"`
	Hmac     string    `json:"hmac,omitempty" binding:"required" db:"hmac"`
}

type FileChunk struct {
	Id          int64 `json:"-" db:"id"`
	FileId      int64 `json:"-" db:"fileId"`
	S3Path      string    `json:"-" db:"s3Path"`
	Size        int64    `json:"size,omitempty" db:"_size"`
	UploadUrl   string    `json:"uploadUrl,omitempty"`
	DownloadUrl string `json:"downloadUrl,omitempty"`
	Order       int `json:"order" db:"_order"`
	Created     time.Time `json:"-" db:"created"`
}

type FileDisplay struct {
	FileHeader FileHeader `json:"fileHeader"`
	FileChunks []FileChunk `json:"chunks"`
}

