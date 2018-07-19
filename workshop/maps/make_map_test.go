package maps

import (
	"reflect"
	"testing"
)

var expected map[string]int

func createExpected() {
	expected = map[string]int{
		"alpha":   5,
		"beta":    4,
		"gamma":   5,
		"delta":   5,
		"epsilon": 7,
	}
}

func TestMakeMap(t *testing.T) {
	m := MakeMap()
	if !reflect.DeepEqual(m, expected) {
		t.Errorf("Expected %v got %v", m, expected)
	}
}

func TestMapLiteral(t *testing.T) {
	m := MapLiteral()
	if !reflect.DeepEqual(m, expected) {
		t.Errorf("Expected %v got %v", m, expected)
	}
}

func TestMain(m *testing.M) {
	createExpected()
	m.Run()
}
