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

	//Add Count block is written here assuming there no DB involved and only will be keep count of names during single runtime.
	//count.GlobalCount[name]++
	/*if val, ok := count.GlobalCount[name]; ok {
		_ = val
		count.GlobalCount[name]++
	} else {
		count.GlobalCount[name] = 1
	}*/

	s.WriteHeader(http.StatusOK)
	fmt.Fprintln(s, "Hello,", name, "!")

}

//
//func GetCounts(s http.ResponseWriter, r *http.Request) {
//	jsonString, _ := json.Marshal(count.GlobalCount)

//	s.WriteHeader(http.StatusOK)
//	fmt.Println(jsonString)
//}

//
func DeleteCounts(s http.ResponseWriter, r *http.Request) {

}

//
func NotFound(s http.ResponseWriter, r *http.Request, status int) {
	s.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprintln(s, "404")
	}
}
