package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"github.com/menkveldj/nafue-api/services"
	"github.com/menkveldj/nafue-api/models"
	"github.com/menkveldj/nafue-api/config"
)

type FileController struct {
	fileService services.IFileService
}

func (self *FileController) Init(routes *config.Routes) {
	self.fileService = services.GetFileService()
	routes.Public.GET("/files/:file", self.getFile)
	routes.Public.POST("/files", self.addFile)
}

/**
 * @api {get} /api/files/:shortUrl Get File
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
func (self *FileController) getFile(c *gin.Context) {

	fileKey := c.Param("file")

	fileDisplay := models.FileHeader{
		ShortUrl: fileKey,
	}

	err := self.fileService.GetFile(&fileDisplay)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusNotFound, gin.H{"message": "Files doesn't exist."})
		return
	}

	c.JSON(http.StatusOK, fileDisplay)
}

/**
 * @api {post} /api/files/ Add File
 * @apiName addFile
 * @apiGroup Files
 *
 * @apiParam {String} aData Data to verify encryption.
 * @apiParam {String} iv Initialization vector for encryption (Base 64 Encoded)
 * @apiParam {String} salt Salt for password (Base 64 Encoded)
 *
 * @apiSuccess {String} aData Data to verify encryption.
 * @apiSuccess {String} uploadUrl S3 pre-signed put request for encrypted file
 * @apiSuccess {String} iv Initialization vector for encryption (Base 64 Encoded)
 * @apiSuccess {String} salt Salt for password (Base 64 Encoded)
 * @apiSuccess {String} shortUrl filename used in request
 *
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 200 OK
 *     {
 *	"aData": "U3BhcnRpY3Vz"
 *	"uploadUrl": "https://s3.amazonaws.com/files.nafue.com/files/4b18adb8-1c6d-45b6-40ff-4c925e67ea23?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAJA7EK4C2E54DYW5Q%2F20160131%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20160131T052532Z&X-Amz-Expires=900&X-Amz-SignedHeaders=host&X-Amz-Signature=58a15149474b217c9a6dc5d3be97252848fdef0edbf7412bdc8fc306e3df7a95"
 *	"iv": "b6GzN+S86GvJ9xUPJG6Rvw=="
 *	"salt": "fNJglILcuQM="
 *	"shortUrl": "eb4d9a3f-2799-4b69-6120-babee1029463"
 *     }
 *
 * @apiError FileNotFound The file cannot be saved.
 *
 * @apiErrorExample Error-Response:
 *     HTTP/1.1 404 Not Found
 *     {
 *       "message": "File cannot be saved."
 *     }
 */
func (self *FileController) addFile(c *gin.Context) {
	// read req body
	var fileDisplay models.FileHeader
	err := c.BindJSON(&fileDisplay)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Body is poorly formated"})
		return
	}

	// add file to db
	err = self.fileService.AddFile(&fileDisplay)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error saving file"})
		return
	}

	c.JSON(http.StatusOK, fileDisplay)
}
