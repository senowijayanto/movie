package repository

import (
	"backend/model/domain"
	"context"
	"database/sql"
)

type IMovieRepository interface {
	ListById(ctx context.Context, tx *sql.Tx, movieId int) (res domain.Movie, err error)
}

type MovieRepository struct{}

func NewMovieRepository() IMovieRepository {
	return &MovieRepository{}
}

func (repo *MovieRepository) ListById(ctx context.Context, tx *sql.Tx, movieId int) (res domain.Movie, err error) {
	return
}
