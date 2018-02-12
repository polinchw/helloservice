package count

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestGetCount(t *testing.T) {
	req, err := http.NewRequest("GET", "/counts", nil)
	if err != nil {
		t.Fatalf("could not created request: %v", err)
	}

	res := httptest.NewRecorder()
	GetCounts(res, req)
	for k, v := range count {
		if res.Body.String() == ("{" + "\"name\":" + "\"" + k + "\"," + "\"count\": " + strconv.Itoa(v) + "}") {
			t.Error(res.Body.String())
		}
	}

}
