package models

import "time"

type FileHeader struct {
	Id         int64       `json:"-" db:"id"`
	TTL        int64       `json:"ttl,omitempty" db:"ttl"`
	Created    time.Time `json:"-" db:"created"`
	ShortUrl   string    `json:"shortUrl,omitempty" db:"shortUrl"`
	AsyncKey   string    `json:"asyncKey,omitempty" db:"asyncKey"`
	ChunkCount int64    `json:"chunkCount,omitempty"`
}

type FileChunk struct {
	Id          int64 `json:"-" db:"id"`
	FileId      int64 `json:"-" db:"fileId"`
	S3Path      string    `json:"-" db:"s3Path"`
	UploadUrl   string    `json:"uploadUrl,omitempty"`
	DownloadUrl string `json:"downloadUrl,omitempty"`
	Order       int `json:"order" db:"_order"`
	Created     time.Time `json:"-" db:"created"`
}

type FileDisplay struct {
	FileHeader FileHeader `json:"fileHeader"`
	FileChunks map[string]FileChunk `json:"fileChunks"`
}

