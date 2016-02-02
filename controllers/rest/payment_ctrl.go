package rest

import (
	"github.com/gin-gonic/gin"
	"nafue/services"
	"net/http"
	"nafue/models/display"
)

type PaymentController struct {
	paymentService services.IPaymentService
}

func (self *PaymentController) Init(r *gin.Engine) {
	self.paymentService = services.GetPaymentService()
	r.GET("/api/clientToken", self.getClientToken)
}

/**
 * @api {get} /api/clientToken Get Client Payment Token
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
	paymentTokenDisplay := display.PaymentTokenDisplay {}

	err := self.paymentService.GetClientToken(&paymentTokenDisplay)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	c.JSON(http.StatusOK, paymentTokenDisplay)
	return
}
