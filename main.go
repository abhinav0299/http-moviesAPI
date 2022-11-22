package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

//	type test struct {
//		Fname    string `json:"first"`
//		Lname    string `json:"title"`
//		Enroll   int
//		Password string   `json:"-"`
//		Tags     []string `json:"tags,omitempty"`
//	}
//
//	type Course struct {
//		CourseId    string  `json:"courseid"`
//		CourseName  string  `json:"coursename"`
//		CoursePrice int     `json:"price"`
//		Author      *Author `json:"author"`
//	}
//
//	type Author struct {
//		FullName string `json:"fullname"`
//		Website  string `json:"website"`
//	}
type MovieModel struct {
	Code   int64  `json:"code"`
	Status string `json:"status"`
	Data   *Data  `json:"data"`
}
type Data struct {
	Movie *Movie `json:"movie"`
}

type Movie struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Genre    string  `json:"genre"`
	Rating   float64 `json:"rating"`
	Plot     string  `json:"plot"`
	Released bool    `json:"released"`
}

type Ping struct {
	Value string `json:"value"`
}

var counter int = 0
var movies []Movie

// helper file
func (c *Movie) isEmpty() bool {
	return c.Name == ""
}

func main() {

	//n := mux.NewRouter()
	//n.HandleFunc("/ping", sendPing)
	//n.NotFoundHandler = http.HandlerFunc(notFoundHandler)
	r := mux.NewRouter()
	r.HandleFunc("/", serverHome).Methods("GET")
	r.HandleFunc("/movies", getAllMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovieById).Methods("GET")
	r.HandleFunc("/movies", addOneMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateOneMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
	//listen to port
	log.Fatal(http.ListenAndServe(":4000", r))
	//log.Fatal(http.ListenAndServe(":4000", n))
}
func serverHome(w http.ResponseWriter, r *http.Request) {
	fmt.Println("server home")
	w.Write([]byte("<h1>Welcome to API by LearnCodeOnline</h1>"))
}

//	func sendPing(w http.ResponseWriter, r *http.Request) {
//		if r.Method == "GET" {
//			ping := Ping{
//				"ping",
//			}
//			counter += 1
//			w.Header().Set("Content - Type", "application/json")
//			w.Header().Set("X-Req-Count", strconv.Itoa(counter))
//			json.NewEncoder(w).Encode(&ping)
//		} else {
//			w.WriteHeader(http.StatusForbidden)
//			json.NewEncoder(w).Encode("Method not allowed")
//		}
//	}
//
//	func notFoundHandler(w http.ResponseWriter, r *http.Request) {
//		w.Header().Set("Content - Type", "application/json")
//		w.WriteHeader(http.StatusNotFound)
//		json.NewEncoder(w).Encode("Not found")
//	}
func getAllMovies(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get all movies")
	w.Header().Set("Content - Type", "application/json")
	json.NewEncoder(w).Encode(&movies)
}
func getMovieById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get all movies by id")
	w.Header().Set("Content - Type", "application/json")
	params := mux.Vars(r)
	for _, movie := range movies {
		movieId, _ := strconv.Atoi(params["id"])
		if movie.ID == movieId {
			responseOutput := MovieModel{
				Code:   200,
				Status: "SUCCESS",
				Data:   &Data{Movie: &movie},
			}
			json.NewEncoder(w).Encode(&responseOutput)
			return
		}
	}
	json.NewEncoder(w).Encode("No movie Found with given id")
	return
}
func addOneMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Println("add movies")
	w.Header().Set("Content - Type", "application/json")
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please Send data")
	}
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	if movie.isEmpty() {
		json.NewEncoder(w).Encode("No data inside json")
	}
	for _, val := range movies {
		if movie.Name == val.Name {
			json.NewEncoder(w).Encode("same movie already there")
			return
		}
	}
	movie.ID = 1
	//rand.Seed(time.Now().UnixNano())
	//movie.ID = rand.Intn(1000)
	movies = append(movies, movie)
	responseOutput := MovieModel{
		Code:   200,
		Status: "SUCCESS",
		Data:   &Data{Movie: &movie},
	}
	json.NewEncoder(w).Encode(&responseOutput)
	return
}
func updateOneMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Println("update movies")
	w.Header().Set("Content - Type", "application/json")
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send data")
	}
	params := mux.Vars(r)
	for ind, movie := range movies {
		movieId, _ := strconv.Atoi(params["id"])
		if movie.ID == movieId {
			var newMovie Movie
			_ = json.NewDecoder(r.Body).Decode(&newMovie)
			newMovie.ID = movieId
			newMovie.Name = movie.Name
			newMovie.Genre = movie.Genre
			movies = append(movies[:ind], movies[ind+1:]...)
			movies = append(movies, newMovie)
			responseOutput := MovieModel{
				Code:   200,
				Status: "SUCCESS",
				Data:   &Data{Movie: &newMovie},
			}
			json.NewEncoder(w).Encode(&responseOutput)
			return
		}
	}

}
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delete courses")
	w.Header().Set("Content - Type", "application/json")
	params := mux.Vars(r)
	for ind, movie := range movies {
		movieId, _ := strconv.Atoi(params["id"])
		if movie.ID == movieId {
			movies = append(movies[:ind], movies[ind+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode("deleted")
	return
}

//func EncodeJson() {
//	details := []test{
//		{"abhinav", "singh", 112, "dada", []string{
//			"sde1", "sde2"}}, {"purvi", "singh", 113, "dadu", nil},
//	}
//	finalJson, err := json.MarshalIndent(details, "", "\t")
//	if err != nil {
//		panic(err)
//	}
//	fmt.Printf("%s\n", finalJson)
//}

//func DecodeJson() {
//	jsonDataFromWeb := []byte(`
//	{
//		"first": "abhinav",
//		"title": "singh",
//		"Enroll": 112,
//		"tags": ["sde1", "sde2"]
//    }
//    `)
//	var details test
//	isValid := json.Valid(jsonDataFromWeb)
//	if isValid {
//		json.Unmarshal(jsonDataFromWeb, &details)
//		fmt.Printf("%#v", details)
//	} else {
//		fmt.Println("Json not valid")
//	}
//	var keyVal map[string]interface{}
//	json.Unmarshal(jsonDataFromWeb, &keyVal)
//	fmt.Printf("%#v", keyVal)
//
//}
