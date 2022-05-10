package controller

import (
	"backend/helper"
	"backend/model/web"
	"backend/service"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"
)

type IMovieController interface {
	ListById(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	ListAll(w http.ResponseWriter, r *http.Request, params httprouter.Params)
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

func (c *MovieController) ListAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var limit, offset int

	// Filter/Search
	filterStr := r.URL.Query().Get("filter")
	filter := ""
	if filterStr != "" {
		filter = filterStr
	}

	// SortBy
	sortStr := r.URL.Query().Get("sortBy")
	if sortStr == "" {
		sortStr = "average.desc"
	}
	splitSort := strings.Split(sortStr, ".")
	field, order := splitSort[0], splitSort[1]

	// Set Limit for pagination
	limitStr := r.URL.Query().Get("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil && limitStr != "" {
		panic(err)
	}

	// Set Offset for pagination
	offsetStr := r.URL.Query().Get("offset")
	offset, err = strconv.Atoi(offsetStr)
	if err != nil && offsetStr != "" {
		panic(err)
	}

	paginationParam := helper.FetchParam{
		Limit:  limit,
		Offset: offset,
		Sort:   fmt.Sprintf("%s %s", field, strings.ToUpper(order)),
		Filter: filter,
	}

	fetchParam := helper.FetchParam{
		Sort:   fmt.Sprintf("%s %s", field, strings.ToUpper(order)),
		Filter: filter,
	}

	var res interface{}

	if limit > 0 || offset > 0 {
		res = c.MovieService.ListAllWithPagination(r.Context(), paginationParam)
	} else {
		res = c.MovieService.ListAll(r.Context(), fetchParam)
	}

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   res,
	}

	helper.WriteToResponseBody(w, webResponse)
}
