package tests

import (
	"testing"
	"net/http"
	"bytes"
)

const baseUrl  = "http://localhost:3000/"

func TestFindAllMovies(t *testing.T)  {
	req, _ := http.NewRequest("GET", baseUrl + "movies", nil)
	res, err := http.DefaultClient.Do(req)
	t.Log(res)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 200 {
		t.Error(`Wrong status code`)
	}
}

func TestCreateMovie(t *testing.T)  {
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
