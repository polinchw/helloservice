package service

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
