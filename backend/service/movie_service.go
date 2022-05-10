package service

import (
	"backend/exception"
	"backend/helper"
	"backend/model/domain"
	"backend/repository"
	"context"
	"database/sql"
)

type IMovieService interface {
	ListById(ctx context.Context, movieId int) domain.Movie
	ListAll(ctx context.Context, params helper.FetchParam) (res []domain.MovieRatingResponse)
	ListAllWithPagination(ctx context.Context, params helper.FetchParam) (res []domain.MovieRatingResponse)
}

type MovieService struct {
	MovieRepository repository.IMovieRepository
	DB              *sql.DB
}

func NewMovieService(movieRepository repository.IMovieRepository, DB *sql.DB) IMovieService {
	return &MovieService{
		MovieRepository: movieRepository,
		DB:              DB,
	}
}

func (service *MovieService) ListById(ctx context.Context, movieId int) domain.Movie {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	movie, err := service.MovieRepository.ListById(ctx, tx, movieId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return movie
}

func (service *MovieService) ListAll(ctx context.Context, params helper.FetchParam) (res []domain.MovieRatingResponse) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	res = service.MovieRepository.ListAll(ctx, tx, params)

	return
}

func (service *MovieService) ListAllWithPagination(ctx context.Context, params helper.FetchParam) (res []domain.MovieRatingResponse) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	res = service.MovieRepository.ListAllWithPagination(ctx, tx, params)

	return
}
