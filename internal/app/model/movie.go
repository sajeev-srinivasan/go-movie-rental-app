package model

type Movie struct {
	Id     string `json:"id"`
	Title  string `json:"title"`
	Year   int    `json:"year"`
	Genre  string `json:"genre"`
	Actors string `json:"actors"`
}
