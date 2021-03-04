package bracketsvalid20

import "testing"

func TestValid(t *testing.T) {
	tests := []struct {
		s   string
		rlt bool
	}{
		{"()[]{}", true},
		{"(]", false},
		{"([)]", false},
		{"{[]}", true},
	}
	for _, aa := range tests {
		got := isValid(aa.s)
		if got != aa.rlt {
			t.Fatalf("error:%v, got:%v,  want:%v", aa.s, got, aa.rlt)
		}
	}
}
