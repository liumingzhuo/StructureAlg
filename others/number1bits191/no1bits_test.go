package number1bits191

import "testing"

func TestNumber(t *testing.T) {
	got := hammingWeight(00000000000000000000000010000000)
	want := 1
	if got != want {
		t.Fatalf("got:%v, but want %v", got, want)
	}
}
