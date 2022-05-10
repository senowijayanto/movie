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
	ListAll(ctx context.Context, tx *sql.Tx, params helper.FetchParam) (res []domain.MovieRatingResponse)
	ListAllWithPagination(ctx context.Context, tx *sql.Tx, params helper.FetchParam) (res []domain.MovieRatingResponse)
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

func (repo *MovieRepository) ListAll(ctx context.Context, tx *sql.Tx, params helper.FetchParam) (res []domain.MovieRatingResponse) {
	sqlSubs := "SELECT AVG(r.rating) FROM ratings r WHERE r.movie_id = m.movie_id GROUP BY movie_id"
	avgQuery := fmt.Sprintf("(%s) AS average", sqlSubs)
	queryBuilder := sq.Select("m.title", "m.release_date", avgQuery).From("movies m")
	if params.Filter != "" {
		queryBuilder = queryBuilder.Where("m.title LIKE '%" + params.Filter + "%'")
	}
	queryBuilder = queryBuilder.OrderBy(params.Sort)
	query, args, err := queryBuilder.ToSql()
	helper.PanicIfError(err)

	rows, err := tx.QueryContext(ctx, query, args...)
	helper.PanicIfError(err)
	defer rows.Close()

	res = []domain.MovieRatingResponse{}
	for rows.Next() {
		movie := domain.MovieRating{}
		movieResponse := domain.MovieRatingResponse{}
		err := rows.Scan(&movie.Title, &movie.ReleaseDate, &movie.Rating)
		if movie.Rating.Valid {
			movieResponse.Rating = movie.Rating.Float64
		}
		movieResponse.Title = movie.Title
		movieResponse.ReleaseDate = movie.ReleaseDate
		helper.PanicIfError(err)
		res = append(res, movieResponse)
	}

	return res
}

func (repo *MovieRepository) ListAllWithPagination(ctx context.Context, tx *sql.Tx, params helper.FetchParam) (res []domain.MovieRatingResponse) {
	sqlSubs := "SELECT AVG(r.rating) FROM ratings r WHERE r.movie_id = m.movie_id GROUP BY movie_id"
	avgQuery := fmt.Sprintf("(%s) AS average", sqlSubs)
	queryBuilder := sq.Select("m.title", "m.release_date", avgQuery).From("movies m")
	if params.Filter != "" {
		queryBuilder = queryBuilder.Where("m.title LIKE '%" + params.Filter + "%'")
	}
	queryBuilder = queryBuilder.OrderBy(params.Sort)

	if params.Limit > 0 {
		queryBuilder = queryBuilder.Limit(uint64(params.Limit))
	}

	if params.Offset > 0 {
		queryBuilder = queryBuilder.Offset(uint64(params.Offset))
	}

	query, args, err := queryBuilder.ToSql()
	helper.PanicIfError(err)

	rows, err := tx.QueryContext(ctx, query, args...)
	helper.PanicIfError(err)
	defer rows.Close()

	res = []domain.MovieRatingResponse{}
	for rows.Next() {
		movie := domain.MovieRating{}
		movieResponse := domain.MovieRatingResponse{}
		err := rows.Scan(&movie.Title, &movie.ReleaseDate, &movie.Rating)
		if movie.Rating.Valid {
			movieResponse.Rating = movie.Rating.Float64
		}
		movieResponse.Title = movie.Title
		movieResponse.ReleaseDate = movie.ReleaseDate
		helper.PanicIfError(err)
		res = append(res, movieResponse)
	}

	return res
}
