package services

import (
	"github.com/lionelbarrow/braintree-go"
	"log"
	"github.com/menkveldj/nafue-api/config"
	"github.com/menkveldj/nafue-api/models"
)

type IPaymentService interface {
	GetClientToken(*models.Payment) error
	ProcessNonce(*models.Payment) error
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
	bt.MerchantAccount()
	return &PaymentService{bt}
}

func (self *PaymentService) GetClientToken(paymentTokenDisplay *models.Payment) error {
	log.Println("merch account: " + config.BtMerchActId)
	token, err := self.bt.ClientToken().Generate()
	if err != nil {
		log.Println("ERROR getting token: ", err.Error())
		return err
	}
	paymentTokenDisplay.Token = token
	return nil
}

func (self *PaymentService) ProcessNonce(paymentNonceDisplay *models.Payment) error {
	// marshal decimal
	dAmount := &braintree.Decimal{}
	err := dAmount.UnmarshalText([]byte(paymentNonceDisplay.Amount))
	if err != nil {
		log.Println("ERROR: failed to create decimal", err)
	}

	result, err := self.bt.Transaction().Create(&braintree.Transaction{
		Type:   "sale",
		Amount: dAmount,
		MerchantAccountId: config.BtMerchActId,
		PaymentMethodNonce: paymentNonceDisplay.Nonce,
	})
	if err != nil {
		log.Println("ERROR: Processing Nonce: ", err.Error())
		return err
	}

	log.Println("Nonce Result: ", result)
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
