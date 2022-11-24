package Movie

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/abhinav0299/moviesApi/internal/models"
)

type Store struct {
	db *sql.DB
}

func New(db *sql.DB) *Store {
	return &Store{db}
}

func (s *Store) GetBookById(movieId int) (models.MovieModel, error) {
	res, err := s.db.QueryContext(context.Background(), "select * from moviesTable where id = ?", movieId)
	var responseOutput models.MovieModel
	if err != nil {
		responseOutput = models.MovieModel{
			Error: err,
		}
		return responseOutput, err
	} else {
		var resMovie models.Movie
		val := res.Next()
		if val {
			err = res.Scan(&resMovie.ID, &resMovie.Name, &resMovie.Genre, &resMovie.Rating, &resMovie.Plot, &resMovie.Released)
			if err != nil {
				responseOutput = models.MovieModel{
					Error: err,
				}
				return responseOutput, err
			}
		} else {
			responseOutput = models.MovieModel{
				Error: err,
			}
			return responseOutput, err
		}
		responseOutput = models.MovieModel{
			Code:   200,
			Status: "SUCCESS",
			Data:   &models.Data{Movie: &resMovie},
		}
	}
	return responseOutput, nil
}
func (s *Store) AddOneMovie(movie models.Movie) (models.MovieModel, error) {
	var responseOutput models.MovieModel
	res, er := s.db.ExecContext(context.Background(), `INSERT INTO moviesTable(name,genre,rating,plot,released) values (?,?,?,?,?)`, movie.Name, movie.Genre, movie.Rating, movie.Plot, movie.Released)
	if er != nil {
		responseOutput = models.MovieModel{
			Error: er,
		}
		return responseOutput, er
	}
	resp, err := res.LastInsertId()
	if err != nil {
		responseOutput = models.MovieModel{
			Error: err,
		}
		return responseOutput, er
	}
	movie.ID = int(resp)
	responseOutput = models.MovieModel{
		Code:   200,
		Status: "SUCCESS",
		Data:   &models.Data{Movie: &movie},
	}
	return responseOutput, nil
}
func (s *Store) UpdateMovieById(id int, newMovie models.Movie) (models.MovieModel, error) {
	var responseOutput models.MovieModel
	var resMovie models.Movie

	_, err := s.db.ExecContext(context.Background(), "UPDATE moviesTable SET rating=?,plot=?,released=? WHERE id=?", newMovie.Rating, newMovie.Plot, newMovie.Released, id)
	if err != nil {
		responseOutput = models.MovieModel{
			Error: err,
		}
		return responseOutput, err
	}
	data, err := s.db.QueryContext(context.Background(), "SELECT * FROM moviesTable WHERE id=?", id)
	if err != nil {
		fmt.Println("inside err")
		responseOutput = models.MovieModel{
			Error: err,
		}
		return responseOutput, err
	}
	val := data.Next()
	if val {
		errs := data.Scan(&resMovie.ID, &resMovie.Name, &resMovie.Genre, &resMovie.Rating, &resMovie.Plot, &resMovie.Released)
		if errs != nil {
			responseOutput = models.MovieModel{
				Error: errs,
			}
			return responseOutput, errs
		}
	}
	responseOutput = models.MovieModel{
		Code:   200,
		Status: "SUCCESS",
		Data:   &models.Data{Movie: &resMovie},
	}
	return responseOutput, nil
}
