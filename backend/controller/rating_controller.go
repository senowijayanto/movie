package controller

import (
	"backend/helper"
	"backend/model/domain"
	"backend/model/web"
	"backend/service"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type IRatingController interface {
	SaveRating(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}

type RatingController struct {
	RatingService service.IRatingService
}

func NewRatingController(ratingService service.IRatingService) IRatingController {
	return &RatingController{
		RatingService: ratingService,
	}
}

func (c *RatingController) SaveRating(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	ratingCreateRequest := domain.Rating{}
	helper.ReadFromRequestBody(r, &ratingCreateRequest)

	ratingResponse := c.RatingService.SaveRating(r.Context(), ratingCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   ratingResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}
