package repository

import (
	"backend/helper"
	"backend/model/domain"
	"context"
	"database/sql"
)

type IActorRepository interface {
	ListAll(ctx context.Context, tx *sql.Tx) []domain.Actor
}

type ActorRepository struct{}

func NewActorRepository() IActorRepository {
	return &ActorRepository{}
}

func (repo *ActorRepository) ListAll(ctx context.Context, tx *sql.Tx) []domain.Actor {
	SQL := "SELECT actor_id, name FROM actors ORDER BY actor_id ASC"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var actors []domain.Actor
	for rows.Next() {
		actor := domain.Actor{}
		err := rows.Scan(&actor.Id, &actor.Name)
		helper.PanicIfError(err)
		actors = append(actors, actor)
	}
	return actors
}
