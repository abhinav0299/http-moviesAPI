package main

import (
	"database/sql"
	"errors"
	"fmt"
	httpLayer "github.com/abhinav0299/moviesApi/internal/http/Movie"
	serviceLayer "github.com/abhinav0299/moviesApi/internal/services/Movie"
	storeLayer "github.com/abhinav0299/moviesApi/internal/stores/Movie"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func getDbObject() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:my-secret-pw@tcp(localhost:3306)/movies")
	if err != nil {
		return nil, errors.New(err.Error())
	} else {
		fmt.Println("db connected")
	}
	return db, nil
}
func main() {
	r := mux.NewRouter()
	db, err := getDbObject()
	if err != nil {
		panic(err)
	}
	storeHandler := storeLayer.New(db)
	serviceHandler := serviceLayer.New(storeHandler)
	httpHandler := httpLayer.New(serviceHandler)
	r.HandleFunc("/movies/{id}", httpHandler.GetMovieById).Methods("GET")
	r.HandleFunc("/movies", httpHandler.AddOneMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", httpHandler.UpdateMovie).Methods("PUT")
	//r.HandleFunc("/movies/{id}", httpHandler.DeleteMovie).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":4000", r))
}
