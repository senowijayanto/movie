package domain

type Rating struct {
	RatingId int `json:"rating_id"`
	MovieId  int `json:"movie_id"`
	Rating   int `json:"rating"`
}
