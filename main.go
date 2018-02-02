package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/hello:{name}", GetHelloName).Methods("GET")
	log.Fatal(http.ListenAndServe(":12345", router))
}
