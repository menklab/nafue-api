package rest

import (
	"github.com/gin-gonic/gin"
	"log"
	"nafue/models/display"
	"nafue/services"
	"net/http"
)

type PaymentController struct {
	paymentService services.IPaymentService
}

func (self *PaymentController) Init(r *gin.Engine) {
	self.paymentService = services.GetPaymentService()
	r.GET("/api/payment", self.getClientToken)
	r.POST("/api/payment", self.processNonce)
}

/**
 * @api {get} /api/payment Get Client Payment Token
 * @apiName getClientToken
 * @apiGroup Payment
 *
 * @apiSuccessExample Success-Response:
 *     	HTTP/1.1 200 OK
 * @apiSuccess {String} token client transaction token
 *	{
 *	"token": "U3BhcnRpY3Vz"
 *     	}
 */
func (self *PaymentController) getClientToken(c *gin.Context) {
	paymentTokenDisplay := display.PaymentTokenDisplay{}

	err := self.paymentService.GetClientToken(&paymentTokenDisplay)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	c.JSON(http.StatusOK, paymentTokenDisplay)
	return
}

/**
 * @api {post} /api/payment Process Payment Nonce
 * @apiName processNonce
 * @apiGroup Payment
 *
 * @apiParam {String} amount Dollar amount to process payment for.
 * @apiParam {String} nonce One-Time nonce authorization for charge.
 *
 * @apiSuccessExample Success-Response:
 *     	HTTP/1.1 200 OK
 */
func (self *PaymentController) processNonce(c *gin.Context) {
	var paymentNonceDisplay display.PaymentNonceDisplay
	err := c.BindJSON(&paymentNonceDisplay)
	if err != nil {
		log.Println("couldn't marshel paymentNonceDisplay: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Body is poorly formated"})
		return
	}

	err = self.paymentService.ProcessNonce(&paymentNonceDisplay)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	c.String(http.StatusOK, "ok")
	return
}
