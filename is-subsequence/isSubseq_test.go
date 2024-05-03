package issubsequence

import "testing"

func TestIsSubsequence(t *testing.T) {
	got := issubsequence("ace", "abcde")
	want := true
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
