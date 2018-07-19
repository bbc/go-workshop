package slices

import (
	"fmt"
	"reflect"
	"testing"
)

// StrToArray ...
func TestStrToSlice(t *testing.T) {
	expected := []rune{'f', 'i', 'l', 'e', 'h', 'o', 'u', 'n', 'd'}
	actual := StrToSlice("filehound")
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf(fmt.Sprintf("Expected %v got %v", expected, actual))
	}
}
