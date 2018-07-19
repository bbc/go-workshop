package functions

import (
	"fmt"
	"testing"
)

// ConvertMobile ..
func TestConvertMobile(t *testing.T) {
	numbers := map[string]string{
		"07704567899":  "+447704567899",
		"07704 567899": "+447704567899",
		"0770456789":   "",
	}

	for k, v := range numbers {
		number, err := ConvertMobile(k)
		if len(v) > 0 && number != v {
			t.Fatalf("Expected %q got %q", v, number)

		} else if err != nil {
			mesg := fmt.Sprintf("Invalid format for: %s", k)
			if err.Error() != mesg {
				t.Fatalf("Expected %q got %q", mesg, err.Error())
			}
		}
	}
}
