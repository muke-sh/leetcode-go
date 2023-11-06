package sorted_squaredarray

import (
	"reflect"
	"testing"
)

func TestSortedSquaredArray(t *testing.T) {
	input := []int{-7, -5, -4, 3, 6, 8, 9}
	want := []int{9, 16, 25, 36, 49, 64, 81}

	got := SortedSquaredArray(input)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %#v, but want %#v", got, want)
	}

}
