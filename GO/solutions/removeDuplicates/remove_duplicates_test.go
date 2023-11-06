package removeduplicates

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	testCases := []struct {
		input  []int
		output int
	}{
		{input: []int{1, 1, 2}, output: 2},
		{input: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, output: 5},
	}

	for _, testCase := range testCases {
		got := removeDuplicates(testCase.input)
		want := testCase.output

		if got != want {
			t.Errorf("got %d, but want %d", got, want)
		}

	}
}
