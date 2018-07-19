package misc

import (
	"io/ioutil"
	"log"
	"strings"
	"testing"
)

// DO NOT MODIFY THESE TESTS - REPORT ANY FAULTS TO DA CHEF :-)

func TestWithCarry(t *testing.T) {
	x := 123
	y := 321
	actual := Sum(x, y)
	if actual != x+y {
		t.Errorf("Expected %d got %d", x+y, actual)
	}
}

func TestWithoutCarry(t *testing.T) {
	x := 4
	y := 8
	actual := Sum(x, y)
	if actual != x+y {
		t.Errorf("Expected %d got %d", x+y, actual)
	}
}

func TestVerifyNoArthemicOperators(t *testing.T) {
	content, err := ioutil.ReadFile("./sum.go")
	if err != nil {
		log.Fatal(err)
	}

	if strings.ContainsAny(string(content), "+-") {
		t.Fatalf("Invalid operators used in solution")
	}
}
