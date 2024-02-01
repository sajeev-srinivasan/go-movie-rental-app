package service

import (
	"movie-rental-app/internal/app/model"
	"movie-rental-app/internal/app/repository"
)

type MovieService interface {
	GetMovies() ([]model.Movie, error)
}

type movieService struct {
	movieRepository repository.MovieRepository
}

func NewMovieService(movieRepository repository.MovieRepository) MovieService {
	return &movieService{movieRepository: movieRepository}
}

func (m movieService) GetMovies() ([]model.Movie, error) {
	movies, err := m.movieRepository.GetMovies()
	if err != nil {
		return []model.Movie{}, err
	}
	return movies, nil
}
