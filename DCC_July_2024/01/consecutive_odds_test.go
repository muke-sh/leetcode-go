package consecutive_odd

import "testing"

func TestThreeConsecutiveOdds(t *testing.T) {

	testCases := []struct {
		input []int
		want  bool
	}{
		{input: []int{2, 6, 4, 1}, want: false},
		{input: []int{1, 2, 34, 3, 4, 5, 7, 23, 12}, want: true},
	}

	for _, tc := range testCases {

		got := threeConsecutiveOdds(tc.input)
		if got != tc.want {
			t.Errorf("got %v, but want %v for input: %#v", got, tc.want, tc.input)
		}
	}

}
