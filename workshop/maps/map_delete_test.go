package maps

import (
	"reflect"
	"testing"
)

func TestMapDelete(t *testing.T) {
	expected := 16
	n := map[string]int{
		"alpha":   5,
		"beta":    4,
		"epsilon": 7,
	}
	if uniqSum, m := MapDelete(); uniqSum != expected {
		t.Errorf("Expected %d got %d", expected, uniqSum)
		if !reflect.DeepEqual(m, n) {
			t.Errorf("Expected %v got %v", n, m)
		}
	}
}
