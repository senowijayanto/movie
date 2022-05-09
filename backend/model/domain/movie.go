package domain

type Movie struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Release     int    `json:"release"`
	Description string `json:"description"`
}
