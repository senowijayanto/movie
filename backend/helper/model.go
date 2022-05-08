package helper

import (
	"backend/model/domain"
	"backend/model/web"
)

func ActorResponse(actor domain.Actor) web.ActorResponse {
	return web.ActorResponse{
		Id:   actor.Id,
		Name: actor.Name,
	}
}

func ActorsResponse(actors []domain.Actor) []web.ActorResponse {
	var actorsResponse []web.ActorResponse

	for _, actor := range actors {
		actorsResponse = append(actorsResponse, ActorResponse(actor))
	}
	return actorsResponse
}
