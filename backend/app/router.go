package app

import (
	"backend/controller"
	"backend/exception"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(actorController controller.IActorController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/actors", actorController.ListAll)

	router.PanicHandler = exception.ErrorHandler

	return router
}
