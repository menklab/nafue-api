package display

import (
	"fmt"
)

type PaymentNonceDisplay struct {
	Amount string `json:"amount" binding:"required"`
	Nonce  string `json:"nonce" binding:"required"`
}

func (self *PaymentNonceDisplay) ToString() string {
	return fmt.Sprintf(
		"{Amount: %v, Nonce: %v}",
		self.Amount,
		self.Nonce,
	)
}
