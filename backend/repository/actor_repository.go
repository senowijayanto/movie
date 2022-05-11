package repository

import (
	"backend/helper"
	"backend/model/domain"
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"
)

type IActorRepository interface {
	ListAll(ctx context.Context, tx *sql.Tx, params helper.FetchParam) (res []domain.Actor)
}

type ActorRepository struct{}

func NewActorRepository() IActorRepository {
	return &ActorRepository{}
}

func (repo *ActorRepository) ListAll(ctx context.Context, tx *sql.Tx, params helper.FetchParam) (res []domain.Actor) {
	queryBuilder := sq.Select("actor_id", "name").From("actors").OrderBy("actor_id")

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

	res = []domain.Actor{}
	for rows.Next() {
		actor := domain.Actor{}
		err := rows.Scan(&actor.Id, &actor.Name)
		helper.PanicIfError(err)
		res = append(res, actor)
	}
	return res
}
