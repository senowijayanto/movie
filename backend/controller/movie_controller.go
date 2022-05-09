package controller

import (
	"backend/helper"
	"backend/model/web"
	"backend/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type IMovieController interface {
	ListById(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}

type MovieController struct {
	MovieService service.IMovieService
}

func NewMovieController(movieService service.IMovieService) IMovieController {
	return &MovieController{
		MovieService: movieService,
	}
}

func (c *MovieController) ListById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	movieId := params.ByName("movieId")
	id, err := strconv.Atoi(movieId)
	helper.PanicIfError(err)

	res := c.MovieService.ListById(r.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   res,
	}

	helper.WriteToResponseBody(w, webResponse)
}
