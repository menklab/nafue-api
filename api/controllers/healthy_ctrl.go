package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/menkveldj/nafue-api/config"
)

type HealthyController struct{}

func (self *HealthyController) Init(routes *config.Routes) {
	routes.Public.GET("/api/healthy", self.healthy)
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
	return
}
