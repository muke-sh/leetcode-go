package twosum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {

	testCases := []struct {
		nums   []int
		target int
		want   []int
	}{
		{nums: []int{3, 2, 4}, target: 6, want: []int{1, 2}},
		{nums: []int{2, 7, 11, 15}, target: 9, want: []int{0, 1}},
		{nums: []int{3, 3}, target: 6, want: []int{0, 1}},
	}

	for _, test := range testCases {
		got := TwoSum(test.nums, test.target)
		want := test.want
		assert(t, got, want, test.nums)
	}

}

func TestTwoSumHashMap(t *testing.T) {

	testCases := []struct {
		nums   []int
		target int
		want   []int
	}{
		{nums: []int{3, 2, 4}, target: 6, want: []int{1, 2}},
		{nums: []int{2, 7, 11, 15}, target: 9, want: []int{0, 1}},
		{nums: []int{3, 3}, target: 6, want: []int{0, 1}},
	}

	for _, test := range testCases {
		got := TwoSumHashMap(test.nums, test.target)
		want := test.want
		assert(t, got, want, test.nums)
	}

}

func assert(t testing.TB, got, want, input []int) {

	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v but want %v with input %v", got, want, input)
	}
}
