package rest

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"io/ioutil"
	"sparticus/domain/models"

	"encoding/json"
	"log"
	"errors"
)

type FileController struct{}

func (self *FileController) Init(router *httprouter.Router) {
		router.POST("/api/file", ResponseHandler(self.AddFile))
}

func (self *FileController) AddFile(w http.ResponseWriter, req *http.Request, params httprouter.Params) (interface{}, httpStatus) {
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return "", serverError(err)
	}

	var file models.File
	err = json.Unmarshal(data, &file)
	if err != nil {
		log.Println(err.Error())
		return "", serverError(errors.New("Invalid Request"))
	}

//	fileId, err :=

	return nil, statusOk(http.StatusOK)
}


