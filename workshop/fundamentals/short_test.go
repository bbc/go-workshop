package fundamentals

import "testing"

func TestShortVars(t *testing.T) {
	var boiling, freezing, absolute float64 = ShortVars()
	if boiling != 100.0 {
		t.Errorf("Expected %f got %f", 100.0, boiling)
	}

	if freezing != 0.0 {
		t.Errorf("Expected %f got %f", 0.0, freezing)
	}

	if absolute != -273.150000 {
		t.Errorf("Expected %f got %f", -273.150000, absolute)
	}
}
