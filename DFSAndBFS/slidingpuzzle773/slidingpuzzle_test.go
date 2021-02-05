package slidingpuzzle773

import "testing"

func TestSlid(t *testing.T) {

	got := slidingPuzzle([][]int{{1, 2, 3}, {4, 0, 5}})
	want := 1
	if got != want {
		t.Fatalf("got: %v, but want %v", got, want)
	}
}
