package addtwonumbers

import (
	"reflect"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	testCases := []struct {
		List1 *ListNode
		List2 *ListNode
		want  *ListNode
	}{
		{List1: CreateLinkedList([]int{9, 9, 9, 9, 9, 9, 9}),
			List2: CreateLinkedList([]int{9, 9, 9, 9}),
			want:  CreateLinkedList([]int{8, 9, 9, 9, 0, 0, 0, 1})},
	}

	for _, test := range testCases {
		got := AddTwoNumbers(test.List1, test.List2)
		assert(t, got, test.want)
	}
}

func assert(t testing.TB, got, want *ListNode) {

	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %s but want %s", got, want)
	}
}
