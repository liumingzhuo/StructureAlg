package minwindowsub76

import (
	"testing"
)

func TestMin(t *testing.T) {
	got := minWindow("a", "aa")
	want := "a"
	if got != want {
		t.Errorf("error want :%s  but got :%s", want, got)
	}
}
