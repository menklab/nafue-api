package rest

import (
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
	"sparticus/domain/models"

	"encoding/json"
	"errors"
	"log"
	"sparticus/services"
)

type FileController struct {
	fileService services.IFileService
}

func (self *FileController) Init(router *httprouter.Router) {
	self.fileService = services.GetFileService()
	router.POST("/api/file", ResponseHandler(self.AddFile))
}

func (self *FileController) AddFile(w http.ResponseWriter, req *http.Request, params httprouter.Params) (interface{}, httpStatus) {
	// read req body
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return "", serverError(err)
	}
	// set to json
	file := models.File{}
	err = json.Unmarshal(data, &file)
	if err != nil {
		log.Println(err.Error())
		return "", serverError(errors.New("Invalid Request"))
	}

	// add file to db
	fileId, err := self.fileService.AddFile()
	if err != nil {
		log.Println(err.Error())
		return "", serverError(errors.New("Error adding file."))
	}
	file.Id = fileId

	// return ok
	return file, statusOk(http.StatusOK)
}
