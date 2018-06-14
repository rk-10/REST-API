package main

import (
	"net/http"
	"log"
	"fmt"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

type Movie struct {
	ID bson.ObjectId	`bson:"id" json:"id"`
	Name string			`bson:"name" json:"name"`
	CoverImage string	`bson:"cover_image" json:"cover_image"`
	Description string	`bson:"description" json:"description"`
}

func AllMoviesEndPoint(w http.ResponseWriter, r *http.Request)  {
	fmt.Println(w, "Not implemented yet")
}

func FindMovieEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func CreateMovieEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func UpdateMovieEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func DeleteMovieEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}


func main()  {
	r := mux.NewRouter()
	r.HandleFunc("/movies", AllMoviesEndPoint).Methods("GET")
	r.HandleFunc("/movies", CreateMovieEndPoint).Methods("POST")
	r.HandleFunc("/movies", UpdateMovieEndPoint).Methods("PUT")
	r.HandleFunc("/movies", DeleteMovieEndPoint).Methods("DELETE")
	r.HandleFunc("/movies/{id}", FindMovieEndpoint).Methods("GET")
	http.ListenAndServe(":3000", r)
	//http.Handle("/", r)
	log.Fatal("Error")
}