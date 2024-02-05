package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"movie-rental-app/internal/app/model"
)

type MovieRepository interface {
	GetMovies(year string, genre string, actors string) ([]model.Movie, error)
	GetAllMovies() ([]model.Movie, error)
	GetMovie(id string) (model.Movie, error)
}

type movieRepository struct {
	*sql.DB
}

func NewMovieRepository(DB *sql.DB) MovieRepository {
	return &movieRepository{DB: DB}
}

func (m movieRepository) GetAllMovies() ([]model.Movie, error) {
	var movies []model.Movie
	rows, err := m.DB.Query("select * from movies")
	if err != nil {
		return []model.Movie{}, errors.New("unable to fetch data" + err.Error())
	}
	return executeQuery(rows, movies)
}

func (m movieRepository) GetMovies(year string, genre string, actors string) ([]model.Movie, error) {
	var movies []model.Movie
	q := fmt.Sprintf("select * from movies where year = %s or genre = '%s' or actors = '%s'", year, genre, actors)
	fmt.Println("q-->", q)
	rows, err := m.DB.Query(q)
	if err != nil {
		return []model.Movie{}, errors.New("unable to fetch data" + err.Error())
	}
	return executeQuery(rows, movies)
}

func (m movieRepository) GetMovie(id string) (model.Movie, error) {
	var movie model.Movie
	row := m.DB.QueryRow("select * from movies where id=$1", id)
	if err := row.Scan(&movie.Id, &movie.Title, &movie.Year, &movie.Genre, &movie.Actors); err != nil {
		fmt.Println("error in fetching data ow", err.Error())
		return model.Movie{}, err
	}
	return movie, nil
}

func executeQuery(rows *sql.Rows, movies []model.Movie) ([]model.Movie, error) {
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println("error while closing rows ->", err.Error())
		}
	}(rows)

	for rows.Next() {
		var movie model.Movie
		if err := rows.Scan(&movie.Id, &movie.Title, &movie.Year, &movie.Genre, &movie.Actors); err != nil {
			fmt.Println("error in fetching data ow", err.Error())
			return movies, err
		}
		movies = append(movies, movie)
	}
	return movies, nil
}
