package controller

import (
	"backend/helper"
	"backend/model/web"
	"backend/service"
	"net/http"
	"strconv"

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
	var limit, offset int

	limitStr := r.URL.Query().Get("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil && limitStr != "" {
		panic(err)
	}

	offsetStr := r.URL.Query().Get("offset")
	offset, err = strconv.Atoi(offsetStr)
	if err != nil && offsetStr != "" {
		panic(err)
	}

	fetchParam := helper.FetchParam{
		Limit:  limit,
		Offset: offset,
	}

	actorsResponse := c.ActorService.ListAll(r.Context(), fetchParam)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   actorsResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}
