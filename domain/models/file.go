package models

import "time"

type File struct {
	Id     int    `json:"id"`
	S3Path string `json:"s3Path"`
	TTL    int    `json:"ttl"`
	ShortUrl string `json:"shortUrl"`
	Created time.Time `json:"lastlogin"`
}
