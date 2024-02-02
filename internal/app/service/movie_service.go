package service

import (
	"movie-rental-app/internal/app/model"
	"movie-rental-app/internal/app/repository"
)

type MovieService interface {
	GetMovies(year string, genre string, actors string) ([]model.Movie, error)
}

type movieService struct {
	movieRepository repository.MovieRepository
}

func NewMovieService(movieRepository repository.MovieRepository) MovieService {
	return &movieService{movieRepository: movieRepository}
}

func (m movieService) GetMovies(year string, genre string, actors string) ([]model.Movie, error) {
	var (
		movies []model.Movie
		err    error
	)
	if year != "" || genre != "" || actors != "" {
		movies, err = m.movieRepository.GetMovies(year, genre, actors)
	} else {
		movies, err = m.movieRepository.GetAllMovies()
	}
	if err != nil {
		return []model.Movie{}, err
	}
	return movies, nil
}
