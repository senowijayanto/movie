package domain

import "database/sql"

type Movie struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	ReleaseDate string `json:"release_date"`
	Description string `json:"description"`
}

type MovieRating struct {
	Title       string `json:"title"`
	ReleaseDate string `json:"release_date"`
	Rating      sql.NullFloat64
}

type MovieRatingResponse struct {
	Title       string  `json:"title"`
	ReleaseDate string  `json:"release_date"`
	Rating      float64 `json:"rating"`
}
