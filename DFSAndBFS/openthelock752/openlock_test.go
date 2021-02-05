package openthelock752

import (
	"fmt"
	"testing"
)

type TestStruct struct {
	Deapends []string
	Target   string
	step     int
}

func TestOpenLock(t *testing.T) {
	tests := []TestStruct{
		{[]string{"0201", "0101", "0102", "1212", "2002"}, "0202", 6},
		{[]string{"8888"}, "0009", 1},
		{[]string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"}, "8888", -1},
	}
	for _, test := range tests {
		got := openLock(test.Deapends, test.Target)
		if got != test.step {
			fmt.Printf("got:%v  but want %v\n", got, test.step)
		}
	}
}
