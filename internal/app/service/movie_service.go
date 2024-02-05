package service

import (
	"fmt"
	"movie-rental-app/internal/app/constants"
	"movie-rental-app/internal/app/model"
	"movie-rental-app/internal/app/repository"
)

type MovieService interface {
	GetMovies(year string, genre string, actors string) ([]model.Movie, error)
	GetMovie(id string) (model.Movie, error)
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

func (m movieService) GetMovie(id string) (model.Movie, error) {
	movie, err := m.movieRepository.GetMovie(id)
	if err != nil {
		fmt.Println("----> ", err.Error())
		if err.Error() == "sql: no rows in result set" {
			return model.Movie{}, constants.ErrNoSuchMovie
		}
	}
	return movie, nil
}
