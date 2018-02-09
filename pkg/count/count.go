package count

import (
	"bytes"
	"fmt"
	"net/http"
	"strconv"
)

/*Assignment hints that reviewer is expecting an array data structure to be used,
however MAP is much more memory efficient - therfore I have used MAP over ARRAY*/
var count map[string]int

//AddCount block is written here assuming there is no DB involved and only will be keep count of names during single runtime.
func AddCount(name string) {
	if count == nil {
		count = make(map[string]int)
	}
	if _, exists := count[name]; exists {
		count[name]++
	} else {
		count[name] = 1
	}
}

//GetCounts will returns a JSON response with the counts of how many times a name has been called.
func GetCounts(s http.ResponseWriter, r *http.Request) {

	var buffer bytes.Buffer

	buffer.WriteString("[")
	i := 0
	for k, v := range count {
		buffer.WriteString("{" + "\"name\":" + "\"" + k + "\"," + "\"count\": " + strconv.Itoa(v) + "}")
		if i != len(count)-1 {
			buffer.WriteString(",")
		}
		i++
	}

	buffer.WriteString("]")
	s.WriteHeader(http.StatusOK)
	fmt.Fprintln(s, buffer.String())
}

//DeleteCounts will resets the data so there are no counts
func DeleteCounts(s http.ResponseWriter, r *http.Request) {
	count = nil
}
