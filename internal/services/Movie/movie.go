package Movie

import (
	"github.com/abhinav0299/moviesApi/internal/models"
	"github.com/abhinav0299/moviesApi/internal/stores"
)

type Service struct {
	store stores.MovieStorer
}

func New(s stores.MovieStorer) *Service {
	return &Service{store: s}
}

func (s *Service) AddOneMovie(movie models.Movie) (models.MovieModel, error) {
	m, err := s.store.AddOneMovie(movie)
	return m, err
}
func (s *Service) UpdateMovieById(id int, movie models.Movie) (models.MovieModel, error) {
	m, err := s.store.UpdateMovieById(id, movie)
	return m, err
}
func (s *Service) GetBookById(movieId int) (models.MovieModel, error) {
	m, err := s.store.GetBookById(movieId)
	return m, err
}
