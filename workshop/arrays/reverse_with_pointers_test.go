package arrays

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	letters := [4]string{"a", "b", "c", "d"}
	expected := [4]string{"d", "c", "b", "a"}

	Reverse(letters)
	if !reflect.DeepEqual(expected, letters) {
		t.Fatal(fmt.Sprintf("Expected %v got %v", expected, letters))
	}
}
