package subsets78

import (
	"log"
	"reflect"
	"testing"
)

func TestSub(t *testing.T) {
	got := subsets([]int{1, 2, 3})
	want := [][]int{
		{},
		{1},
		{2},
		{1, 2},
		{3},
		{1, 3},
		{2, 3},
		{1, 2, 3},
	}
	if !reflect.DeepEqual(got, want) {
		log.Fatalf("got:%v, but want:%v", got, want)
	}

}
