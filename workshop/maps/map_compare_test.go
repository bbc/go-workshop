package maps

import "testing"

func TestCompareWhereEqual(t *testing.T) {
	m := map[string]string{
		"a": "b",
		"b": "c",
		"c": "d",
	}

	n := map[string]string{
		"a": "b",
		"b": "c",
		"c": "d",
	}

	if !MapCompare(m, n) {
		t.Errorf("Maps expected to be equal")
	}
}

func TestCompareWhereNotEqual(t *testing.T) {
	m := map[string]string{
		"a": "b",
		"b": "c",
		"c": "d",
	}

	n := map[string]string{
		"z": "b",
		"b": "c",
		"c": "d",
	}

	if MapCompare(m, n) {
		t.Errorf("Maps expected to be equal")
	}
}
