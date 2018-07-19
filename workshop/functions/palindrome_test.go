package functions

import (
	"fmt"
	"testing"
)

func TestReturnsTrueWhenIsPalindrome(t *testing.T) {
	actual := IsPalindrome("aabaa")
	if !actual {
		t.Fatal(fmt.Sprintf("Expected %v, got %v", true, actual))
	}
}

func TestReturnsFalseWhenNotAPalindrome(t *testing.T) {
	actual := IsPalindrome("aabcaa")
	if actual {
		t.Fatal(fmt.Sprintf("Expected %v, got %v", false, actual))
	}
}
