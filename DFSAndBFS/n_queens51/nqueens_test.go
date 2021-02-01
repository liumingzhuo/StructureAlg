package n_queens51

import (
	"reflect"
	"testing"
)

func TestNQ(t *testing.T) {
	got := solveNQueens(4)
	want := [][]string{{".Q..", "...Q", "Q...", "..Q."}, {"..Q.", "Q...", "...Q", ".Q.."}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("excepted;	got:%v  want:%v ", got, want)
	}

}
