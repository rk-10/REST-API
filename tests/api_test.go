package tests

import (
	"testing"
	"net/http"
	"net/http/httptest"

	"github.com/gorilla/mux"
)


func getMovies(t *testing.T)  {
	req, _ := http.NewRequest("GET", "/", nil)
	response := executeRequest(req)

	if response.Code != http.StatusOK {
		t.Error("Expected status code to be 200")
	}

}

func executeRequest(req *http.Request)  *httptest.ResponseRecorder{
	router := mux.NewRouter()
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)
	return  recorder
}