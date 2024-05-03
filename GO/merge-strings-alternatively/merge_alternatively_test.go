package mergestringsalternatively

import "testing"

func TestMergeAlternately(t *testing.T) {
	testCases := []struct {
		input  []string
		output string
	}{
		{
			input:  []string{"abc", "pqr"},
			output: "apbqcr",
		},
		{
			input:  []string{"ab", "pqrs"},
			output: "apbqrs",
		},
		{
			input:  []string{"abcd", "pq"},
			output: "apbqcd",
		},
		{
			input:  []string{"cf", "eee"},
			output: "cefee",
		},
	}

	for _, testCase := range testCases {
		got := mergeAlternately(testCase.input[0], testCase.input[1])
		want := testCase.output

		if got != want {
			t.Errorf("got %q, but want %q", got, want)
		}
	}
}
