package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/somyagarg94/helloservice/pkg"
)

var addr = flag.String("addr", ":12345", "The address to listen on for HTTP requests.")

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/hello:{name}", pkg.GetHelloName).Methods("GET")
	router.HandleFunc("/health", pkg.HealthChecker).Methods("GET")
	router.HandleFunc("/counts", pkg.GetCounts).Methods("GET")
	router.HandleFunc("/counts", pkg.DeleteCounts).Methods("DELETE")
	log.Fatal(http.ListenAndServe(*addr, router))
}
