package slices

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	letters := []string{"a", "b", "c", "d"}
	expected := []string{"d", "c", "b", "a"}

	Reverse(letters)
	if !reflect.DeepEqual(expected, letters) {
		t.Fatal(fmt.Sprintf("Expected %v got %v", expected, letters))
	}
}
