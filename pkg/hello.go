package pkg

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//GetHelloName returns the hello names
func GetHelloName(s http.ResponseWriter, r *http.Request) {
	log.Println("Responsing to /hello:name request")

	vars := mux.Vars(r)
	name := vars["name"]

	s.WriteHeader(http.StatusOK)
	fmt.Fprintln(s, "Hello,", name, "!")

}

//
func GetCounts(s http.ResponseWriter, r *http.Request) {

	//var count int
	//count++
	//s.Header().Set("X-Call-Count", fmt.Fprintln("%d", count))

}

//
func DeleteCounts(s http.ResponseWriter, r *http.Request) {

}
