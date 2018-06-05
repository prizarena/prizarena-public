package prizarena

import "testing"

func TestApiError_Error(t *testing.T) {
	err := apiError{Code: "code1", Message: "msg1"}
	if err.Error() != "code1: msg1" {
		t.Fatalf("Unexpeceted result")
	}
}
