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
	"bytes"
)

var dao = MoviesDAO{}


func AllMoviesEndPoint(w http.ResponseWriter, r *http.Request)  {
	movies, err := dao.FindAll()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Bad request. Please check all parameters."))
		fmt.Println("Error in params")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
	w.WriteHeader(http.StatusOK)
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
		fmt.Println("Error in params")
		return
	}
	movie.ID = bson.NewObjectId()

	if err := dao.Insert(movie); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("400 - Could not store data to db"))
		fmt.Println("Data to db could not be stored")
		return
	}

	w.WriteHeader(http.StatusOK)

	var buffer bytes.Buffer
	buffer.WriteString(`{Response: Success}`)
	json.NewEncoder(w).Encode(buffer.String())

	fmt.Println("Data stored to db")
}

func UpdateMovieEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var movie Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Bad request. Please check all parameters."))
		fmt.Println("Error in params")
		return
	}

	if err := dao.Update(movie); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Something went wrong with the server"))
		fmt.Println("Error in updating")
		return
	}

	w.WriteHeader(http.StatusOK)

	var buffer bytes.Buffer
	buffer.WriteString(`{Response: Success}`)
	json.NewEncoder(w).Encode(buffer.String())

	fmt.Println("Movie updated to db")
}

func DeleteMovieEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func init()  {
	dao.Server = "localhost:27017"
	dao.Database = "test"
	dao.Connect()
}


func main()  {
	r := mux.NewRouter()
	r.HandleFunc("/movies", AllMoviesEndPoint).Methods("GET")
	r.HandleFunc("/movies", CreateMovieEndPoint).Methods("POST")
	r.HandleFunc("/movies", UpdateMovieEndPoint).Methods("PUT")
	r.HandleFunc("/movies", DeleteMovieEndPoint).Methods("DELETE")
	r.HandleFunc("/movies/{id}", FindMovieEndpoint).Methods("GET")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal("Error", err)
	}
	fmt.Println("Server is up and running on port 3000")
}