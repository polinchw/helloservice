package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/somyagarg94/helloservice/pkg"
	"github.com/somyagarg94/helloservice/pkg/count"
	"github.com/somyagarg94/helloservice/pkg/health"
	"github.com/somyagarg94/helloservice/pkg/hello"
)

var addr = flag.String("addr", ":12345", "The address to listen on for HTTP requests.")

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/hello:{name}", hello.GetHelloName).Methods("GET")
	router.HandleFunc("/health", health.GetHealthChecker).Methods("GET")
	router.HandleFunc("/counts", count.GetCounts).Methods("GET")
	router.HandleFunc("/counts", pkg.DeleteCounts).Methods("DELETE")
	log.Fatal(http.ListenAndServe(*addr, router))
}
