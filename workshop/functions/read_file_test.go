package functions

import (
	"strings"
	"testing"
)

func validateContents(contents []string) bool {
	expectedContents := map[string]int{
		"Lorem ipsum dolor sit amet": 0,
		"LAST": len(contents) - 1,
	}

	for match, index := range expectedContents {
		if !strings.Contains(contents[index], match) {
			return false
		}
	}
	return true
}

func TestContents(t *testing.T) {
	expectedLength := 48
	actual, _ := ReadFile("../fixtures/text.txt")

	if len(actual) != expectedLength {
		t.Fatal("Unexpected file length!")
	}

	if !validateContents(actual) {
		t.Fatal("Unexpected file contents!")
	}
}

func TestInvalidFile(t *testing.T) {
	_, err := ReadFile("../fixtures/invalid.txt")

	if err == nil || !strings.Contains(err.Error(), "no such file") {
		t.Fatal("No error returned for invalid file")
	}
}
