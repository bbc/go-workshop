package slices

import (
	"fmt"
	"testing"
)

func TestReturnsTrueIfTwoWordsAreAnagrams(t *testing.T) {
	anagrams := []string{
		"Desperation", "A rope ends it",
		"Debit card", "bad credit",
		"Dormitory", "dirty room",
	}
	for i := 0; i < len(anagrams); i += 2 {
		anagram := IsAnagram(anagrams[i], anagrams[i+1])
		if !anagram {
			t.Fatalf(fmt.Sprintf("'%s' is an anagram of '%s'. ", anagrams[i], anagrams[i+1]))
		}
	}
}
