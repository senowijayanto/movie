package controller

import (
	"backend/helper"
	"backend/model/web"
	"backend/service"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type IActorController interface {
	ListAll(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}

type ActorController struct {
	ActorService service.IActorService
}

func NewActorController(actorService service.IActorService) IActorController {
	return &ActorController{
		ActorService: actorService,
	}
}

func (c *ActorController) ListAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	actorsResponse := c.ActorService.ListAll(r.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   actorsResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}
