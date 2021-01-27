package jumpgame45

import "testing"

func TestJump(t *testing.T) {
	got := jump([]int{2, 3, 1, 1, 4})
	want := 2
	if got != want {
		t.Fatalf("got %d\n but want %d", got, want)
	}
}
