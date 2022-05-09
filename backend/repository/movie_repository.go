package repository

import (
	"backend/helper"
	"backend/model/domain"
	"context"
	"database/sql"
	"errors"
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

type IMovieRepository interface {
	ListById(ctx context.Context, tx *sql.Tx, movieId int) (res domain.Movie, err error)
}

type MovieRepository struct{}

func NewMovieRepository() IMovieRepository {
	return &MovieRepository{}
}

func (repo *MovieRepository) ListById(ctx context.Context, tx *sql.Tx, movieId int) (res domain.Movie, err error) {
	queryBuilder := sq.Select("movie_id", "title", "release_date", "description").From("movies m").PlaceholderFormat(sq.Question).OrderBy("movie_id")

	whereClause := queryBuilder.Where("movie_id = ?", fmt.Sprintf("%d", movieId))

	query, args, err := whereClause.ToSql()
	helper.PanicIfError(err)
	fmt.Println(query)

	rows, err := tx.QueryContext(ctx, query, args...)
	helper.PanicIfError(err)
	defer rows.Close()

	res = domain.Movie{}
	if rows.Next() {
		err := rows.Scan(&res.Id, &res.Title, &res.ReleaseDate, &res.Description)
		helper.PanicIfError(err)
		return res, nil
	} else {
		return res, errors.New("movie is not found")
	}
}
