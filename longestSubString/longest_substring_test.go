package longestsubstring

import (
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {

	testCases := []struct {
		inputString  string
		wantedLength int
	}{
		{inputString: "abcabcbb", wantedLength: 3},
		{inputString: "bbbbb", wantedLength: 1},
		{inputString: "pwwkew", wantedLength: 3},
		{inputString: "dvdf", wantedLength: 3},
	}

	for _, test := range testCases {
		gotLength := LengthOfLongestSubstring(test.inputString)

		assert(t, gotLength, test.wantedLength, test.inputString)
	}

}

func assert(t testing.TB, got, want int, input string) {
	t.Helper()

	if got != want {
		t.Errorf("got %d but want %d with input %q", got, want, input)
	}
}
