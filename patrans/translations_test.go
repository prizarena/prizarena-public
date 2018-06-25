package patrans

import "testing"

func TestTRANS(t *testing.T) {
	if len(TRANS) == 0 {
		t.Fatalf("TRANS is empty")
	}
}
