package main

import (
	"net/http"
	"log"
	"fmt"
	"github.com/gorilla/mux"
	"encoding/json"
	. "github.com/rk-10/REST-API/models"
	"gopkg.in/mgo.v2/bson"
	. "github.com/rk-10/REST-API/dao"
)

var dao = MoviesDAO{}


func AllMoviesEndPoint(w http.ResponseWriter, r *http.Request)  {
	fmt.Println(w, "Not implemented yet")
}

func FindMovieEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func CreateMovieEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var movie Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Bad request. Please check all parameters."))
		return
	}
	movie.ID = bson.NewObjectId()
	

	fmt.Fprintln(w, "not implemented yet !")
}

func UpdateMovieEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func DeleteMovieEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func init()  {
	dao.Server = ""
	dao.Database = ""
	dao.Connect()
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