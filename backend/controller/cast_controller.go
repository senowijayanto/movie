package controller

import (
	"backend/helper"
	"backend/model/web"
	"backend/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type ICastController interface {
	ListAll(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}

type CastController struct {
	CastService service.ICastService
}

func NewCastController(castService service.ICastService) ICastController {
	return &CastController{
		CastService: castService,
	}
}

func (c *CastController) ListAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
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
	res := c.CastService.ListAll(r.Context(), fetchParam)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   res,
	}

	helper.WriteToResponseBody(w, webResponse)
}
