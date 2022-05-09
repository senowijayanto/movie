package helper

import (
	"backend/model/domain"
)

type FetchParam struct {
	Limit  int
	Offset int
}

func ActorResponse(actor domain.Actor) domain.Actor {
	return domain.Actor{
		Id:   actor.Id,
		Name: actor.Name,
	}
}

func ActorsResponse(actors []domain.Actor) []domain.Actor {
	var actorsResponse []domain.Actor

	for _, actor := range actors {
		actorsResponse = append(actorsResponse, ActorResponse(actor))
	}
	return actorsResponse
}
