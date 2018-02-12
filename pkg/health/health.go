package health

import (
	"encoding/json"
	"log"
	"net/http"
	"runtime"
)

// GetHealthChecker maps results to the expected endpoint response.
func GetHealthChecker(s http.ResponseWriter, r *http.Request) {

	results := make(map[string]float32)

	// get Number of logical CPUs usable by current process.
	numCPU := runtime.NumCPU()
	results["NumOfLogicalCPUsAvailable"] = float32(numCPU)

	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	//Frees is the cumulative count of heap objects freed.
	results["MemFrees"] = float32(memStats.Frees)

	/*
		Other helful information are also available like:
		TotalAlloc 		- TotalAlloc is cumulative bytes allocated for heap objects.
		HeapAlloc 		-  HeapAlloc is bytes of allocated heap objects.
		HeapSys	   	    -  HeapSys is bytes of heap memory obtained from the OS.
		HeapIdle		-  HeapIdle is bytes in idle (unused) spans.
		HeapInuse   	-  HeapInuse is bytes in in-use spans.
		HeapReleased    -  HeapReleased is bytes of physical memory returned to the OS.
		and a lot more options that may help developers and systems engineers to troubleshoot the program are available.
	*/

	res, err := json.Marshal(results)

	if err != nil {
		log.Printf("error: couldn't marshal queue metrics to json")
		s.WriteHeader(http.StatusInternalServerError)
	} else {
		s.Write(res)
	}
}
