package tests

import (
	"testing"
	"net/http"
	"encoding/json"
	"github.com/rk-10/REST-API/models"
	"bytes"
)

const baseUrl  = "http://localhost:3000/"
var movieId string

func TestCreateMovie(t *testing.T) {
	var jsonStr = []byte(`
{
	"name": "Inception",
	"description": "Enigma"
}`)

	req, _ := http.NewRequest("POST", baseUrl + "movies", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)

	defer res.Body.Close()

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 200 {
		t.Error(`Wrong status code`)
	}
}


func TestFindAllMovies(t *testing.T)  {
	req, _ := http.NewRequest("GET", baseUrl + "movies", nil)
	res, err := http.DefaultClient.Do(req)

	var movie []models.Movie;
	json.NewDecoder(res.Body).Decode(&movie)

	movieId = movie[0].ID.Hex()
	t.Log(movieId)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 200 {
		t.Error(`Wrong status code`)
	}


}

//
//func TestUpdateMovie(t *testing.T)  {
//
//}
