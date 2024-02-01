package model

type Response struct {
	Status  string
	Message string
}

type MovieResponse struct {
	Response
	Data []Movie
}
