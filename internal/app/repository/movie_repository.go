package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"movie-rental-app/internal/app/model"
)

type MovieRepository interface {
	GetMovies() ([]model.Movie, error)
}

type movieRepository struct {
	*sql.DB
}

func NewMovieRepository(DB *sql.DB) MovieRepository {
	return &movieRepository{DB: DB}
}

func (m movieRepository) GetMovies() ([]model.Movie, error) {
	var movies []model.Movie
	rows, err := m.DB.Query("select id, title,year,genre,actors from movies")
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println("error while closing rows ->", err.Error())
		}
	}(rows)
	if err != nil {
		return []model.Movie{}, errors.New("unable to fetch data" + err.Error())
	}

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
