package main

import "fmt"

func reverseStr(s string, k int) string {
	b := []byte(s)
	for i := 0; i < len(b); i += 2 * k {
		//k<= b-i<2k
		if i+k <= len(b) {
			reverse(&b, i, i+k-1)
			continue
		}
		//  b-i < k  剩余的b
		reverse(&b, i, len(b)-1)
	}
	return string(b)
}
func reverse(s *[]byte, left, right int) {
	for left < right {
		(*s)[left], (*s)[right] = (*s)[right], (*s)[left]
		left++
		right--
	}

}

func main() {
	got := reverseStr("abcdefg", 2)
	fmt.Println(got)
}
