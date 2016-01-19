package rest

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"sparticus/services"
	"sparticus/models/display"
)

type FileController struct {
	fileService services.IFileService
}

func (self *FileController) Init(r *gin.Engine) {
	self.fileService = services.GetFileService()
	r.POST("/api/files", self.addFile)
	r.GET("/api/files/:file", self.getFile)
}

func (self *FileController) addFile(c *gin.Context) {
	// read req body

	var fileDisplay display.FileDisplay
	err := c.BindJSON(&fileDisplay)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Body is poorly formated"})
		return;
	}



	// add file to db
	err = self.fileService.AddFile(&fileDisplay)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error saving file"})
		return;
	}

	c.JSON(http.StatusOK, fileDisplay)
}


func (self *FileController) getFile(c *gin.Context) {

	fileKey := c.Param("file")

	fileDisplay := display.FileDisplay{
		ShortUrl: fileKey,
	}

	err := self.fileService.GetFile(&fileDisplay)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error retrieving file"})
		return;
	}

	c.JSON(http.StatusOK, fileDisplay)
}