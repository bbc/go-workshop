package arrays

import (
	"reflect"
	"testing"
)

func TestCreate(t *testing.T) {
	expected := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	actual := CreateArray()

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Expected %v got %v", expected, actual)
	}
}
