package repository

import (
	"backend/helper"
	"backend/model/domain"
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"
)

type IRatingRepository interface {
	SaveRating(ctx context.Context, tx *sql.Tx, rating domain.Rating) domain.Rating
}

type RatingRepository struct{}

func NewRatingRepository() IRatingRepository {
	return &RatingRepository{}
}

func (r *RatingRepository) SaveRating(ctx context.Context, tx *sql.Tx, rating domain.Rating) domain.Rating {
	queryBuilder := sq.Insert("ratings").Columns("movie_id", "rating").Values(rating.MovieId, rating.Rating).PlaceholderFormat(sq.Question)
	query, args, err := queryBuilder.ToSql()
	helper.PanicIfError(err)

	result, err := tx.ExecContext(ctx, query, args...)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	rating.RatingId = int(id)
	return rating
}
