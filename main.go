package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var addr = flag.String("addr", ":12345", "The address to listen on for HTTP requests.")

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/hello:{name}", GetHelloName).Methods("GET")
	router.HandleFunc("/health", HealthChecker).Methods("GET")
	router.HandleFunc("/counts", GetCounts).Methods("GET")
	router.HandleFunc("/counts", DeleteCounts).Methods("DELETE")
	log.Fatal(http.ListenAndServe(*addr, router))
}
