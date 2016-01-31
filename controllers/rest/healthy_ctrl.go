package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HealthyController struct{}

func (self *HealthyController) Init(r *gin.Engine) {
	r.GET("/api/healthy", self.healthy)
}

/**
 * @api {get} /api/healthy Health Check
 * @apiName Healthy
 * @apiGroup Basic
 *
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 200 OK
 *
 */
func (self *HealthyController) healthy(c *gin.Context) {
	c.String(http.StatusOK, "ok")
	return;
}
