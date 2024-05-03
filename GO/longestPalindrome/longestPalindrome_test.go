package longestpalindrome

import (
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	testCases := []struct {
		inputString       string
		longestPalindrome string
	}{
		{inputString: "babad", longestPalindrome: "bab"},
		{inputString: "cbbd", longestPalindrome: "bb"},
		{inputString: "a", longestPalindrome: "a"},
		{inputString: "ac", longestPalindrome: "a"},
	}

	for _, test := range testCases {
		got := LongestPalindrome(test.inputString)

		assertStrings(t, got, test.longestPalindrome, test.inputString)
	}
}

func TestReverse(t *testing.T) {

	testCases := []struct {
		inputString string
		want        string
	}{
		{inputString: "Hello", want: "olleH"},
		{inputString: "abcd", want: "dcba"},
	}

	for _, test := range testCases {

		got := Reverse(test.inputString)

		assertStrings(t, got, test.want, test.inputString)
	}

}

func TestIsPalindrome(t *testing.T) {
	testCases := []struct {
		inputString  string
		isPalindrome bool
	}{
		{inputString: "aa", isPalindrome: true},
		{inputString: "aba", isPalindrome: true},
		{inputString: "aab", isPalindrome: false},
	}

	for _, test := range testCases {
		got := IsPalindrome(test.inputString)
		assertBoolean(t, got, test.isPalindrome, test.inputString)
	}
}

func assertStrings(t testing.TB, got, want, input string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q but want %q with input %q", got, want, input)
	}
}

func assertBoolean(t testing.TB, got, want bool, input string) {
	t.Helper()

	if got != want {
		t.Errorf("got %t but want %t with input %q", got, want, input)
	}
}
