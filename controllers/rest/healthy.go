package rest

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type HealthyController struct{}

func (ctrl *HealthyController) Init(router *httprouter.Router) {
		router.GET("/api/healthy", ResponseHandler(ctrl.GetHealthy))
}

func (ctrl *HealthyController) GetHealthy(w http.ResponseWriter, req *http.Request, params httprouter.Params) (interface{}, httpStatus) {
	return nil, StatusOk(http.StatusOK)
}
