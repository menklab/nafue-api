package rest

import (
	"github.com/julienschmidt/httprouter"
)

func Init(router *httprouter.Router) {

	//rest API controllers
	new(HealthyController).Init(router)
}
