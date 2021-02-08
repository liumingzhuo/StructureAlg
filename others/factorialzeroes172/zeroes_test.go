package factorialzeroes172

import "testing"

func TestZeroes(t *testing.T) {
	got := trailingZeroes(25)
	want := 4
	if got != want {
		t.Fatalf("exception: got :%v  , but want :%v", got, want)
	}
}
