package util

import "testing"

func TestEscOffset(t *testing.T) {
	s, i := EscOffset("test")
	if s != "test" && i != 4 {
		t.Error("offset is wrong")
	}
}
