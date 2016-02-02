package display

import "fmt"

type PaymentTokenDisplay struct {
	Token string `json:"token" binding:"required"`
}

func (self *PaymentTokenDisplay) ToString() string {
	return fmt.Sprintf(
		"{Token: %v}",
		self.Token,
	)
}
