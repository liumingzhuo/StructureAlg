package poweroftwo231

import "testing"

func TestPowerTwo(t *testing.T) {
	got := isPowerOfTwo(1)
	want := true
	if got != want {
		t.Fatalf("got:%v but want:%v", got, want)
	}
}
