package main

import (
	"net/http"
	"log"
	"fmt"
	"github.com/gorilla/mux"
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
	. "github.com/rk-10/REST-API/models"
	. "github.com/rk-10/REST-API/dao"
	. "github.com/rk-10/REST-API/config"
)

var dao = MoviesDAO{}
var config = Config{}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func respondWithError(w http.ResponseWriter, code int, msg string)  {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func AllMoviesEndPoint(w http.ResponseWriter, r *http.Request)  {
	movies, err := dao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "400 - Bad request. Please check all parameters.")
		log.Fatal("Error in params")
		return
	}
	respondWithJson(w, http.StatusOK, movies)
}

func FindMovieEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movie, err := dao.FindbyId(params["id"]); if err != nil {
		respondWithError(w, http.StatusBadRequest, "400 - Bad request. Please check all parameters.")
		log.Fatal("Error in params")
		return
	}

	respondWithJson(w, http.StatusOK, movie)
	log.Print(w, "Movie found !")
}

func CreateMovieEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var movie Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		respondWithError(w, http.StatusBadRequest, "400 - Bad request. Please check all parameters.")
		log.Fatal("Error in params")
		return
	}
	movie.ID = bson.NewObjectId()

	if err := dao.Insert(movie); err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not store data to db")
		log.Fatal("Data to db could not be stored")
		return
	}

	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
	log.Print("Data stored to db")
}

func UpdateMovieEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var movie Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		respondWithError(w, http.StatusBadRequest, "400 - Bad request. Please check all parameters.")
		log.Fatal("Error in params")
		return
	}

	if err := dao.Update(movie); err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not store data to db")
		log.Fatal("Error in updating")
		return
	}

	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
	log.Print("Movie updated to db")
}

func DeleteMovieEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var movie Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		respondWithError(w, http.StatusBadRequest, "400 - Bad request. Please check all parameters.")
		log.Fatal("Errror in params")
		return
	}

	if err := dao.Remove(movie); err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not store data to db")
		fmt.Println("Error in updating")
		return
	}

	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
	log.Print("Movie deleted from db")
}

func init()  {
	config.Read()
	dao.Server = config.Server
	dao.Database = config.Database
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
	log.Print("Server is up and running on port 3000")
}