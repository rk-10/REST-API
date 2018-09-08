package tests

import (
	"testing"
	"net/http"
	"encoding/json"
	"github.com/rk-10/REST-API/models"
	"bytes"
)

const baseUrl  = "http://localhost:3000/"
var Movie models.Movie

func TestCreateMovie(t *testing.T) {
	var jsonStr = []byte(`
{
	"name": "Inception",
	"description": "Enigma"
}`)

	req, _ := http.NewRequest("POST", baseUrl + "movies", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 200 {
		t.Error(`Wrong status code`)
	}
}


//This test fails if the above test fails.
func TestFindAllMovies(t *testing.T)  {
	req, _ := http.NewRequest("GET", baseUrl + "movies", nil)
	res, err := http.DefaultClient.Do(req)

	var movie []models.Movie;
	json.NewDecoder(res.Body).Decode(&movie)

	if len(movie) == 0 {
		t.Error(`No movie found in DB`)
	}

	Movie = movie[0]

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 200 {
		t.Error(`Wrong status code`)
	}
}

func TestFindMovieById(t *testing.T)  {

	req, _ := http.NewRequest("GET", baseUrl + "movies/" + Movie.ID.Hex() , nil)
	res, err := http.DefaultClient.Do(req)

	var movie models.Movie;
	json.NewDecoder(res.Body).Decode(&movie)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 200 {
		t.Error(`Wrong status code`)
	}
}


func TestUpdateMovie(t *testing.T)  {
	var jsonStr = []byte(`
{	
	"id": "` + Movie.ID.Hex() + `", 
	"name": "Lala Land",
	"description": "Oscar winning movie!"
}`)

	req, _ := http.NewRequest("PUT", baseUrl + "movies", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 200 {
		t.Error(`Wrong status code`)
	}
}

func TestDeleteMovie(t *testing.T)  {
	var jsonStr = []byte(`
{	
	"id": "` + Movie.ID.Hex() + `", 
	"name": "Lala Land",
	"description": "Oscar winning movie!"
}`)

	req, _ := http.NewRequest("DELETE", baseUrl + "movies", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 200 {
		t.Error(`Wrong status code`)
	}
}
