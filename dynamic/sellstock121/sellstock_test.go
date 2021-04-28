package sellstock121

import "testing"

func TestSellStock(t *testing.T) {
	input := []int{7, 1, 5, 3, 6, 4}
	got := maxProfit(input)
	want := 5
	if got != want {
		t.Errorf("error, want got %d  but got %d", want, got)
	}
}

func TestSellStock2(t *testing.T) {
	input := []int{7, 6, 4, 3, 1}
	got := maxProfit(input)
	want := 0
	if got != want {
		t.Errorf("error, want got %d  but got %d", want, got)
	}
}
