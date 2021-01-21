package NonOverlapIntervals435

import "testing"

func TestNon(t *testing.T) {
	inputNums := [][]int{}
	got := eraseOverlapIntervals(inputNums)
	want := 0
	if got != want {
		t.Fatalf("want %#v  but got %#v", want, got)
	}
}

func TestInter(t *testing.T) {
	type testItem struct {
		arr  [][]int
		want int
	}
	testGroup := []testItem{
		testItem{[][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}, 1},
		testItem{[][]int{{1, 2}, {1, 2}, {1, 2}}, 2},
		testItem{[][]int{{1, 2}, {2, 3}}, 0},
	}
	for _, test := range testGroup {
		got := eraseOverlapIntervals(test.arr)
		if got != test.want {
			t.Fatalf("want %#v  but got %#v", test.want, got)
		}
	}
}
