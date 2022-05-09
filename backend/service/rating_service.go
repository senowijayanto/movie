package service

import (
	"backend/helper"
	"backend/model/domain"
	"backend/repository"
	"context"
	"database/sql"
)

type IRatingService interface {
	SaveRating(ctx context.Context, request domain.Rating) domain.Rating
}

type RatingService struct {
	RatingRepository repository.IRatingRepository
	DB               *sql.DB
}

func NewRatingService(ratingRepository repository.IRatingRepository, DB *sql.DB) IRatingService {
	return &RatingService{
		RatingRepository: ratingRepository,
		DB:               DB,
	}
}

func (service *RatingService) SaveRating(ctx context.Context, request domain.Rating) domain.Rating {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	rating := domain.Rating{
		MovieId: request.MovieId,
		Rating:  request.Rating,
	}

	rating = service.RatingRepository.SaveRating(ctx, tx, rating)

	return rating
}
