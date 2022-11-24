package Movie

import (
	"encoding/json"
	"errors"
	"github.com/abhinav0299/moviesApi/internal/models"
	"github.com/abhinav0299/moviesApi/internal/services"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type Handler struct {
	service services.MovieService
}

func New(s services.MovieService) *Handler {
	return &Handler{service: s}
}

func (h *Handler) GetMovieById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	movieId, _ := strconv.Atoi(params["id"])
	resp, err := h.service.GetBookById(movieId)
	if err != nil {
		responseOutput := models.MovieModel{
			Error: errors.New("No movie Found with given id"),
		}
		json.NewEncoder(w).Encode(&responseOutput)
		return
	}
	json.NewEncoder(w).Encode(&resp)
	return
}
func (h *Handler) AddOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Body == nil {
		responseOutput := models.MovieModel{
			Error: errors.New("No data inside json"),
		}
		json.NewEncoder(w).Encode(&responseOutput)
		return
	}
	var movie models.Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	resp, err := h.service.AddOneMovie(movie)
	if err != nil {
		responseOutput := models.MovieModel{
			Error: errors.New("No movie Found with given id"),
		}
		json.NewEncoder(w).Encode(&responseOutput)
		return
	}
	json.NewEncoder(w).Encode(&resp)
	return
}
func (h *Handler) UpdateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Body == nil {
		responseOutput := models.MovieModel{
			Error: errors.New("No data inside json"),
		}
		json.NewEncoder(w).Encode(&responseOutput)
		return
	}
	params := mux.Vars(r)
	movieId, _ := strconv.Atoi(params["id"])
	var newMovie models.Movie
	_ = json.NewDecoder(r.Body).Decode(&newMovie)
	resp, err := h.service.UpdateMovieById(movieId, newMovie)
	if err != nil {
		responseOutput := models.MovieModel{
			Error: errors.New("No movie Found with given id"),
		}
		json.NewEncoder(w).Encode(&responseOutput)
		return
	}
	json.NewEncoder(w).Encode(&resp)
	return
}

//func (handler *handler) DeleteMovie(w http.ResponseWriter, r *http.Request) {
//	fmt.Println("delete courses")
//	w.Header().Set("Content - Type", "application/json")
//	params := mux.Vars(r)
//	for ind, movie := range movies {
//		movieId, _ := strconv.Atoi(params["id"])
//		if movie.ID == movieId {
//			movies = append(movies[:ind], movies[ind+1:]...)
//			break
//		}
//	}
//	json.NewEncoder(w).Encode("deleted")
//	return
//}
