package CoinChange518

import (
	"reflect"
	"testing"
)

// type testGroup struct {
// 	amount int
// 	coins  []int
// }

func TestCoin(t *testing.T) {
	got := change(5, []int{1, 2, 5})
	want := 4
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want:%#v bug got:%#v", want, got)
	}
}
func TestCoin2(t *testing.T) {
	got := change(3, []int{2})
	want := 0
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want:%#v bug got:%#v", want, got)
	}
}

func TestCoin3(t *testing.T) {
	got := change(10, []int{10})
	want := 1
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want:%#v bug got:%#v", want, got)
	}
}
