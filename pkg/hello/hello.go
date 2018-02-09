package hello

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/somyagarg94/helloservice/pkg/count"
)

//GetHelloName returns the hello <names>.
func GetHelloName(s http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	s.WriteHeader(http.StatusOK)
	count.AddCount(name)
	fmt.Fprintln(s, "Hello,", name, "!")

}
