package service

import (
	"backend/helper"
	"backend/model/domain"
	"backend/repository"
	"context"
	"database/sql"
)

type ICastService interface {
	ListAll(ctx context.Context, params helper.FetchParam) (res []domain.Cast)
}

type CastService struct {
	CastRepository repository.ICastRepository
	DB             *sql.DB
}

func NewCastService(castRepository repository.ICastRepository, DB *sql.DB) ICastService {
	return &CastService{
		CastRepository: castRepository,
		DB:             DB,
	}
}

func (service *CastService) ListAll(ctx context.Context, params helper.FetchParam) (res []domain.Cast) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	res = service.CastRepository.ListAll(ctx, tx, params)

	return
}
