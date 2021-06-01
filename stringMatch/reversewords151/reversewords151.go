/*
 *
 * [151] 翻转字符串里的单词
 *
 */

package main

import (
	"fmt"
)

func reverseWords(s string) string {
	//删除冗余的空格
	slowIndex, fastIndex := 0, 0
	b := []byte(s)
	//头部空格
	for len(b) > 0 && fastIndex < len(b) && b[fastIndex] == ' ' {
		fastIndex++
	}
	for ; fastIndex < len(b); fastIndex++ {
		if fastIndex-1 > 0 && b[fastIndex-1] == b[fastIndex] && b[fastIndex] == ' ' {
			continue
		}
		b[slowIndex] = b[fastIndex]
		slowIndex++
	}
	//尾部空格
	if slowIndex-1 > 0 && b[slowIndex-1] == ' ' {
		b = b[:slowIndex-1]
	} else {
		b = b[:slowIndex]
	}
	//反转整个字符串
	reverse(&b, 0, len(b)-1)
	//反转单个字符  i 单词开始位置，
	i := 0
	for i < len(b) {
		// j 单词结束位置
		j := i
		for ; j < len(b) && b[j] != ' '; j++ {
		}
		reverse(&b, i, j-1)
		i = j
		i++
	}
	return string(b)
}

func reverse(b *[]byte, left, right int) {
	for left < right {
		(*b)[left], (*b)[right] = (*b)[right], (*b)[left]
		left++
		right--
	}
}

func main() {
	got := reverseWords(" the sky  is blue ")
	// want := "blue is sky the"
	fmt.Println(got)

}
