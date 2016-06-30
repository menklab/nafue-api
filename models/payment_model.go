package models


type Payment struct {
	Amount string `json:"amount" binding:"required"`
	Nonce  string `json:"nonce" binding:"required"`
	Token string `json:"token" binding:"required"`

}

