package repository

import (
	"backend/helper"
	"backend/model/domain"
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"
)

type ICastRepository interface {
	ListAll(ctx context.Context, tx *sql.Tx, params helper.FetchParam) (res []domain.Cast)
}

type CastRepository struct{}

func NewCastRepository() ICastRepository {
	return &CastRepository{}
}

func (repo *CastRepository) ListAll(ctx context.Context, tx *sql.Tx, params helper.FetchParam) (res []domain.Cast) {
	queryBuilder := sq.Select("a.name", "m.title").From("movie_casts c").Join("actors a ON c.actor_id = a.actor_id").Join("movies m ON c.movie_id = m.movie_id").PlaceholderFormat(sq.Question).OrderBy("c.actor_id")

	if params.Limit > 0 {
		queryBuilder = queryBuilder.Limit(uint64(params.Limit))
	}

	if params.Offset > 0 {
		queryBuilder = queryBuilder.Offset(uint64(params.Offset))
	}

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return
	}

	rows, err := tx.QueryContext(ctx, query, args...)
	helper.PanicIfError(err)
	defer rows.Close()

	res = []domain.Cast{}
	for rows.Next() {
		cast := domain.Cast{}
		err := rows.Scan(&cast.ActorName, &cast.MovieTitle)
		helper.PanicIfError(err)
		res = append(res, cast)
	}

	return res
}
