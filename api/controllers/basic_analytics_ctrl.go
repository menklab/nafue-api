package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"github.com/menkveldj/nafue-api/services"
	"github.com/menkveldj/nafue-api/models"
	"github.com/menkveldj/nafue-api/config"
)

type BasicAnalyticsController struct {
	basicAnalyticsService services.IBasicAnalyticsService
}

func (self *BasicAnalyticsController) Init(routes *config.Routes) {
	self.basicAnalyticsService = services.GetBasicAnalyticsService()
	routes.Public.GET("/api/basicAnalytics", self.getBasicAnalytics)
}

/**
 * @api {get} /api/basicAnalytics
 * @apiName getBasicAnalytics
 * @apiGroup Analytics
 *
 * @apiSuccess {int} fileCount Total files served since site creation.
 *
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 200 OK
 *     {
 *	"fileCount": 1000
 *     }
 *
 * @apiError FileNotFound The file was not found.
 *
 * @apiErrorExample Error-Response:
 *     HTTP/1.1 404 Not Found
 *     {
 *       "message": "File not found."
 *     }
 */
func (self *BasicAnalyticsController) getBasicAnalytics(c *gin.Context) {

	var basicAnalyticsDisplay models.BasicAnalytics

	err := self.basicAnalyticsService.GetFileCount(&basicAnalyticsDisplay)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusNotFound, gin.H{"message": "Couldn't get basic analytics."})
		return
	}
	c.JSON(http.StatusOK, basicAnalyticsDisplay)
}
