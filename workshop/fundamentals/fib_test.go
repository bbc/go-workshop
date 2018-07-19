package fundamentals

import (
	"testing"
)

func TestNFib(t *testing.T) {
	expected := 55
	n := 10
	nth := Fib(n)
	if nth != expected {
		t.Errorf("Expected %v Got %v", expected, nth)
	}
}
