package health

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthChecker(t *testing.T) {

	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	GetHealthChecker(res, req)
	// Check the result code is what we expect.
	if result := res.Code; result != http.StatusOK {
		t.Errorf("HealthChecker returned wrong result: got %v want %v",
			result, http.StatusOK)
	}

}
