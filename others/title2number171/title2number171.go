/**
171 excel表序列号
1. 将A-Z 表示成1-26
2. 字母* 26^(字母下标距离尾部的长度)
*/
package main

import (
	"fmt"
	"math"
)

func titleToNumber(columnTitle string) int {
	al := []byte(columnTitle)
	n := len(al) - 1
	ans := 0
	for i, v := range al {
		ans += int(float64(v-64) * math.Pow(26, float64(n-i)))
	}
	return ans
}

func main() {
	rlt := titleToNumber("A")
	fmt.Println(rlt)
}
