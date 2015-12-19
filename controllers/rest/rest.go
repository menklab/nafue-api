package rest

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type httpStatus struct {
	err    error
	status int
}

func serverError(err error) httpStatus {
	return httpStatus{err, http.StatusInternalServerError}
}

func statusOk(status int) httpStatus {
	return httpStatus{nil, status}
}

type controllerRoute func(http.ResponseWriter, *http.Request, httprouter.Params) (interface{}, httpStatus)

/* JSON REST utils */

func WriteResponse(w http.ResponseWriter, result interface{}, httpStatus httpStatus) {
	var responseBody string
	w.Header().Set("Content-Type", "application/json")

	if httpStatus.err != nil {
		w.WriteHeader(httpStatus.status)
		jsonBody, _ := json.Marshal(httpStatus.err.Error())
		responseBody = string(jsonBody)
	} else {
		w.WriteHeader(httpStatus.status)
		jsonBody, _ := json.Marshal(result)
		responseBody = string(jsonBody)
	}

	fmt.Fprintf(w, responseBody)
}

func ResponseHandler(r controllerRoute) httprouter.Handle {
	return func(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
		result, httpStatus := r(w, req, p)
		WriteResponse(w, result, httpStatus)
	}
}
