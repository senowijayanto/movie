package repository

import (
	"backend/helper"
	"backend/model/domain"
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"
)

type IActorRepository interface {
	ListAll(ctx context.Context, tx *sql.Tx) (res []domain.Actor)
}

type ActorRepository struct{}

func NewActorRepository() IActorRepository {
	return &ActorRepository{}
}

func (repo *ActorRepository) ListAll(ctx context.Context, tx *sql.Tx) (res []domain.Actor) {
	queryBuilder := sq.Select("actor_id", "name").From("actors").OrderBy("actor_id")
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
