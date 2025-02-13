package arrays

import (
	"testing"
)

func TestContains(t *testing.T) {
	haystack := []string{"apple", "orange", "banana", "mango", "kiwi"}

	got := Contains(haystack, "apple")
	if !got {
		t.Errorf("Contains(%v, \"apple\") = %v; want true", haystack, got)
	}

	got = Contains(haystack, "kiwi")
	if !got {
		t.Errorf("Contains(%v, \"kiwi\") = %v; want true", haystack, got)
	}

	got = Contains(haystack, "cat")
	if got {
		t.Errorf("Contains(%v, \"cat\") = %v; want false", haystack, got)
	}
}
