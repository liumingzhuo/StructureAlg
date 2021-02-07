package singlenumber136

import "testing"

func TestSingle(t *testing.T) {
	got := singleNumber([]int{4, 1, 2, 1, 2})
	want := 4
	if got != want {
		t.Fatalf("got: %v  but want: %v", got, want)
	}
}
