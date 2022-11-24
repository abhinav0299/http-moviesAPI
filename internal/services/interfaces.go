package services

import "github.com/abhinav0299/moviesApi/internal/models"

type MovieService interface {
	AddOneMovie(movie models.Movie) (models.MovieModel, error)
	UpdateMovieById(id int, movie models.Movie) (models.MovieModel, error)
	GetBookById(movieId int) (models.MovieModel, error)
}
