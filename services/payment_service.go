package services

import (
	"github.com/lionelbarrow/braintree-go"
	"log"
	"nafue/config"
	"nafue/models/display"
)

type IPaymentService interface {
	GetClientToken(*display.PaymentTokenDisplay) error
}

type PaymentService struct {
	bt *braintree.Braintree
}

func NewPaymentService() *PaymentService {
	bt := braintree.New(
		getEnv(),
		config.BtMerchId,
		config.BtPubKey,
		config.BtPrivKey,
	)
	return &PaymentService{bt}
}

func (self *PaymentService) GetClientToken(paymentTokenDisplay *display.PaymentTokenDisplay) error {
	token, err := self.bt.ClientToken().Generate()
	if err != nil {
		log.Println("ERROR getting token: ", err.Error())
		return err
	}
	paymentTokenDisplay.Token = token
	return nil
}

func getEnv() braintree.Environment {
	env := braintree.Production

	switch config.BtEnv {
	case "sandbox":
		env = braintree.Sandbox
		break
	case "development":
		env = braintree.Development
	}

	return env
}
