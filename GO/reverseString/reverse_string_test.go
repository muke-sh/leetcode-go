package reversestring

import (
	"reflect"
	"testing"
)

func TestReverseString(t *testing.T) {
	input := []byte{'h', 'e', 'l', 'l', 'o'}
	output := []byte{'o', 'l', 'l', 'e', 'h'}

	ReverseString(input)

	if !reflect.DeepEqual(input, output) {
		t.Errorf("got %v, want %v", input, output)
	}
}
