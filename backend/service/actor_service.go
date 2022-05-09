package service

import (
	"backend/helper"
	"backend/model/domain"
	"backend/repository"
	"context"
	"database/sql"
)

type IActorService interface {
	ListAll(ctx context.Context) []domain.Actor
}

type ActorService struct {
	ActorRepository repository.IActorRepository
	DB              *sql.DB
}

func NewActorService(actorRepository repository.IActorRepository, DB *sql.DB) IActorService {
	return &ActorService{
		ActorRepository: actorRepository,
		DB:              DB,
	}
}

func (service *ActorService) ListAll(ctx context.Context) []domain.Actor {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	actors := service.ActorRepository.ListAll(ctx, tx)

	return actors
}
