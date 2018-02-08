package health

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"runtime"
)

// ErrBadRequest indicates provided parameters did not match specification
var ErrBadRequest = errors.New("bad request")

// GetHealthChecker maps results to the expected endpoint response.
func GetHealthChecker(s http.ResponseWriter, r *http.Request) {
	log.Println("Responsing to /health request")

	results := make(map[string]float32)

	// get number of Goroutines
	numRoutines := runtime.NumGoroutine()
	results["GoRoutines"] = float32(numRoutines)

	// get memory stats
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	// bytes allocated and not yet freed
	results["MemAlloc"] = float32(memStats.Alloc)

	// number of frees
	results["MemFrees"] = float32(memStats.Frees)

	// bytes allocated and not yet freed
	results["MemHeapAlloc"] = float32(memStats.HeapAlloc)

	// total number of allocated objects
	results["MemHeapObjects"] = float32(memStats.HeapObjects)

	// bytes obtained from system
	results["MemHeapSys"] = float32(memStats.HeapSys)

	// number of mallocs
	results["MemMallocs"] = float32(memStats.Mallocs)

	// total number of garbage collections
	results["MemNumGc"] = float32(memStats.NumGC)

	// bytes obtained from system
	results["MemSys"] = float32(memStats.Sys)

	res, err := json.Marshal(results)
	if err != nil {
		log.Printf("error: couldn't marshal queue metrics to json")
		s.WriteHeader(http.StatusInternalServerError)
	} else {
		s.Write(res)
	}
}
