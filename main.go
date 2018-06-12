package main

import (
	"net/http"
	"log"
	"fmt"

	"github.com/gorilla/mux"
)

func Login(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("Not implemented yet")
}

func main()  {
	r := mux.NewRouter()
	r.HandleFunc("/login", Login).Methods("GET")
	http.ListenAndServe(":3000", r)
	//http.Handle("/", r)
	log.Fatal("Error")
}