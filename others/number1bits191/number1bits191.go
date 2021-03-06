package number1bits191

//191. 位1的个数
/*
输入：00000000000000000000000000001011
输出：3
解释：输入的二进制串 00000000000000000000000000001011 中，共有三位为 '1'。

输入：00000000000000000000000010000000
输出：1
解释：输入的二进制串 00000000000000000000000010000000 中，共有一位为 '1'。


输入：11111111111111111111111111111101
输出：31
解释：输入的二进制串 11111111111111111111111111111101 中，共有 31 位为 '1'。

*/
//利用n&(n-1)消除最后一位1
func hammingWeight(num uint32) int {
	var count int
	for num != 0 {
		num = num & (num - 1)
		count++
	}
	return count
}
