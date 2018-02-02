package pkg

import (
	"io"
	"net/http"
)

// HealthChecker checks the health of api and returns the JSON
func HealthChecker(s http.ResponseWriter, r *http.Request) {
	// A very simple health check.
	s.WriteHeader(http.StatusOK)
	s.Header().Set("Content-Type", "application/json")

	// In the future we could report back on the status of our DB, or our cache
	// (e.g. Redis) by performing a simple PING, and include them in the response.
	io.WriteString(s, `{"alive": true}`)
}
