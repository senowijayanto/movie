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
