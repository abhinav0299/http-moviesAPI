package stores

import "github.com/abhinav0299/moviesApi/internal/models"

type MovieStorer interface {
	AddOneMovie(movie models.Movie) (models.MovieModel, error)
	UpdateMovieById(id int, movie models.Movie) (models.MovieModel, error)
	GetBookById(movieId int) (models.MovieModel, error)
}
