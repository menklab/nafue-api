package rest

import (
	"github.com/gin-gonic/gin"
	"log"
	"nafue/models/display"
	"nafue/services"
	"net/http"
)

type BasicAnalyticsController struct {
	basicAnalyticsService services.IBasicAnalyticsService
}

func (self *BasicAnalyticsController) Init(r *gin.Engine) {
	self.basicAnalyticsService = services.GetBasicAnalyticsService()
	r.GET("/api/basicAnalytics", self.getBasicAnalytics)
}

/**
 * @api {get} /api/file/:shortUrl Get File
 * @apiName getFile
 * @apiGroup Files
 *
 * @apiParam {String} shortUrl url used as in place of filename.
 *
 * @apiSuccess {String} aData Data to verify encryption.
 * @apiSuccess {String} downloadUrl S3 pre-signed get request for encrypted data
 * @apiSuccess {String} iv Initialization vector for encryption (Base 64 Encoded)
 * @apiSuccess {String} salt Salt for password (Base 64 Encoded)
 * @apiSuccess {String} shortUrl filename used in request
 *
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 200 OK
 *     {
 *	"aData": "U3BhcnRpY3Vz"
 *	"downloadUrl": "https://s3.amazonaws.com/files.nafue.com/files/4b18adb8-1c6d-45b6-40ff-4c925e67ea23?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAJA7EK4C2E54DYW5Q%2F20160131%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20160131T052532Z&X-Amz-Expires=900&X-Amz-SignedHeaders=host&X-Amz-Signature=58a15149474b217c9a6dc5d3be97252848fdef0edbf7412bdc8fc306e3df7a95"
 *	"iv": "b6GzN+S86GvJ9xUPJG6Rvw=="
 *	"salt": "fNJglILcuQM="
 *	"shortUrl": "eb4d9a3f-2799-4b69-6120-babee1029463"
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

	var basicAnalyticsDisplay display.BasicAnalyticsDisplay

	err := self.basicAnalyticsService.GetFileCount(&basicAnalyticsDisplay)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusNotFound, gin.H{"message": "Couldn't get basic analytics."})
		return
	}
	log.Println("FileCount: ", basicAnalyticsDisplay.ToString())
	c.JSON(http.StatusOK, basicAnalyticsDisplay)
}
