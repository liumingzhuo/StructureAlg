package factorialzeroes172

//172 阶乘后的零
//输入: 3
//输出: 0
//解释: 3! = 6, 尾数中没有零。

//输入: 5
//输出: 1
//解释: 5! = 120, 尾数中有 1 个零.

func trailingZeroes(n int) int {
	count := 0
	for i := n; i/5 > 0; i /= 5 {
		count += i / 5 //只需要检查 n的阶乘中包含多少个5
	}
	return count
}
