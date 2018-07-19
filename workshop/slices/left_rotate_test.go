package slices

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLeftRotate(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6}
	expected := []int{3, 4, 5, 6, 1, 2}

	LeftRotate(arr, 2)
	if !reflect.DeepEqual(arr, expected) {
		t.Fatal(fmt.Sprintf("Expected %v got %v", expected, arr))
	}
}
