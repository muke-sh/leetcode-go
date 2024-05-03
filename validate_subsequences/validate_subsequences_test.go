package validate_subsequences

import "testing"

func TestValidateSubsequences(t *testing.T) {
	testCases := []struct {
		inputs struct {
			arr         []int
			subsequence []int
		}
		output bool
	}{
		{
			inputs: struct {
				arr         []int
				subsequence []int
			}{
				arr:         []int{5, 1, 22, 25, 6, -1, 8, 10},
				subsequence: []int{1, 6, -1, 10},
			},
			output: true,
		},
	}
	_ = testCases

	for _, testCase := range testCases {
		got := validateSubsequence(testCase.inputs.arr, testCase.inputs.subsequence)
		want := testCase.output

		if got != want {
			t.Errorf("got %t, but want %t", got, want)
		}
	}
}
