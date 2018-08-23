package tests

import (
	"testing"
	"net/http"
)


func TestGetAllMovies(t *testing.T)  {
	req, _ := http.NewRequest("GET", "http://localhost:3000/movies", nil)
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 200 {
		t.Error(`Wrong status code`)
	}
}

