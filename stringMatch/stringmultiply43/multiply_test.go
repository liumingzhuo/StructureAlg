package stringmultiply43

import (
	"testing"
)

func TestMultipfy(t *testing.T) {
	got := multiply("2", "3")
	want := "61"
	if got != want {
		t.Fatalf("got: %s  but want: %s", got, want)
	}
}
