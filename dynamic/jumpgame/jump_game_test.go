package jumpgame

import "testing"

func TestJump(t *testing.T) {
	got := canJump([]int{2, 3, 1, 1, 4})
	want := true
	if got != want {
		t.Fatalf("want %#v , but got %#v", want, got)
	}
}
func TestJump2(t *testing.T) {
	got := canJump([]int{3, 2, 1, 0, 4})
	want := false
	if got != want {
		t.Fatalf("want %#v , but got %#v", want, got)
	}
}
