package traprainwater42

import (
	"testing"
)

func TestTrap(t *testing.T) {
	got := trap([]int{4, 2, 0, 3, 2, 5})
	want := 9
	if got != want {
		t.Fatalf("got %d ,but want %d", got, want)
	}
}
