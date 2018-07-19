package testing

import (
	"strings"
)

// PadRight ...
func PadRight(s string, n uint) string {
	size := uint(len(s))
	if size > n {
		return s[:n-1]
	}
	s += strings.Repeat(" ", int(n-size))
	return s
}
