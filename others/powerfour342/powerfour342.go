/**
342. 4的幂
*/
package main

import "fmt"

func isPowerOfFour(n int) bool {
	//判断是2的幂 n&(n-1) ==0
	//判断是4的幂 10101010101010101010101010101010 偶数位为0 奇数位为1
	return n > 0 && n&(n-1) == 0 && n&(0xaaaaaaaa) == 0
}

func main() {
	got := isPowerOfFour(5)
	fmt.Println(got)
}
